package apps

import (
	"bytes"
	"craps/point"
	"embed"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"strconv"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"golang.org/x/exp/rand"
)

var RED, GREEN = color.RGBA{255, 0, 0, 128}, color.RGBA{0, 255, 0, 128}

func makeLabelWithData(title string, filled bool) (bs BS, l *ThemedLabel) {
	bs = NewBS()
	bs.Set(title)
	l = NewThemedLabelWithData(bs)
	l.Alignment = fyne.TextAlignCenter
	if filled {
		l.overlay.FillColor, l.overlay.StrokeWidth = GREEN, 2
	} else {
		l.overlay.StrokeColor, l.overlay.StrokeWidth = GREEN, 2
	}
	return
}

func App1(animationFiles embed.FS, isDark bool) *fyne.Container {

	initialMsg := "Welcome to Simple Craps."
	images := cacheImages(animationFiles)

	pt := point.NewPointTracker()

	ptLabelString, ptLabel := makeLabelWithData(pt.CurState.String(), false)
	pLabelString, pLabel := makeLabelWithData(pt.CurPoint.String(), false)
	playerLabelString, playerLabel := makeLabelWithData(pt.NewPlayer.String(), false)
	ptLabelString.Set(pt.CurState.String())
	pLabelString.Set(pt.CurPoint.String())

	resultString := NewBS()
	resultString.Set(initialMsg)
	result := NewThemedLabelWithData(resultString)
	result.overlay.FillColor = GREEN

	// widgets showing die animation side-by-side
	img := make([]*canvas.Image, 2)

	left, right := 0, 1
	for i := range []int{left, right} {
		img[i] = canvas.NewImageFromImage(images[i][len(images)-1])
		img[i].FillMode, img[i].ScaleMode = canvas.ImageFillOriginal, canvas.ImageScaleFastest
	}

	initialBank := Money(2000)
	savedBank := fyne.CurrentApp().Preferences().Int("bank")
	if savedBank != 0 {
		initialBank = Money(savedBank)
	}
	bank := NewCash(initialBank)
	bankLabel := NewThemedLabelWithData(bank.amtString)
	bankLabel.overlay.StrokeColor, bankLabel.overlay.StrokeWidth = GREEN, 2

	bet := NewBS()
	bet.Set(Money(int64(initialBank) / 2).String())
	betLabel := NewThemedLabelWithData(bet)
	betLabel.overlay.StrokeColor, betLabel.overlay.StrokeWidth = GREEN, 2

	var info *canvas.Text

	textBinding := binding.NewString()
	textBinding.AddListener(binding.NewDataListener(func() {
		newText, _ := textBinding.Get()
		info.Text = newText
		info.Refresh()
	}))

	info = InfoText(textBinding, isDark)

	// the zero based value of the rolled die
	leftDie, rightDie := 0, 0

	var rollButton *widget.Button
	doAnimation := func(tick float32) {
		// there are len(images) to display in 4s, tick will be 0.5 at 2s for instance, which is len(images)/2, so the image # is tick*len(images)
		i := int(tick * float32(len(images[0])-1))
		updateDice(img, left, images, leftDie, i, right, rightDie)
		if tick == 1.0 {
			resultString.Set(initialMsg)
			rollButton.Enable()
			total := leftDie + 1 + rightDie + 1 // Dice values are 1-indexed
			resultText := fmt.Sprintf("You rolled: %d", total)
			pt.SetPoint(total)
			switch pt.CurState {
			case point.COME_OUT_ROLL:
				pt.Reset()
			case point.WIN:
				pt.Reset()
				resultText += ". You Win!"
				bank.AddBetAmt(bet, true)
			case point.LOSE:
				pt.Reset()
				resultText += ". You Lose."
				bank.AddBetAmt(bet, false)
				curBet := ToMoney(bet.GetS())
				if bank.amt <= 0.0 {
					bank.SetAmt(bank.initialBank)
					bet.Set((bank.initialBank / 2).String())
					resultText += " Refreshed bank."
				} else if curBet > bank.amt {
					// limit the bet to the amount of money left in the bank
					bet.Set(bank.String())
				}
			default:
				resultText += ". Roll again."

			}
			ptLabelString.Set(pt.CurState.String())
			pLabelString.Set(pt.CurPoint.String())
			playerLabelString.Set(pt.NewPlayer.String())
			setInfo(textBinding, pt)
			resultString.Set(resultText)
		}
	}

	rand.Seed(uint64(time.Now().UnixNano()))
	rollButton = widget.NewButtonWithIcon("Roll", theme.Icon(theme.IconNameNavigateNext), func() {
		rollButton.Disable()
		b, _ := bet.Get()
		i := ToMoney(b)
		bet.Set(i.String())
		leftDie, rightDie = rand.Intn(6), rand.Intn(6)
		resultString.Set("Rolling...")
		fyne.NewAnimation(1*time.Second, doAnimation).Start()
	})

	keys := make([]fyne.CanvasObject, 0)
	for _, key := range []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "AC", "0", "DEL", "Bet 1/4", "Bet 1/2", "Bet All"} {
		b := widget.NewButton(" "+key+" ", func() { handleKey(key, bet, bank) })
		keys = append(keys, b)
	}

	bankHeading, betHeading := NewThemedLabel("Bank"), NewThemedLabel("Bet")
	bankHeading.Alignment, betHeading.Alignment = fyne.TextAlignCenter, fyne.TextAlignCenter
	bankLabel.Alignment, betLabel.Alignment = fyne.TextAlignCenter, fyne.TextAlignCenter

	result.Alignment = fyne.TextAlignCenter
	result.overlay.StrokeColor, result.overlay.StrokeWidth = GREEN, 2

	bg := canvas.NewRectangle(color.Transparent)
	bg.StrokeWidth, bg.StrokeColor = 2, color.RGBA{128, 128, 128, 128}

	dice := container.NewHBox(
		layout.NewSpacer(),
		img[left],
		img[right],
		layout.NewSpacer(),
	)

	reset := func() {
		pt.Reset()
		initialBank = Money(2000)
		savedBank = int(initialBank)
		bank.Reset(bet, initialBank)
		bank.SetAmt(initialBank)
		bankLabel.Refresh()
	}

	setInfo(textBinding, pt)

	stuff := container.NewVBox(
		dice,
		container.NewGridWithColumns(3, keys...),
		rollButton,
		result.Stack(),
		container.NewGridWithColumns(3, playerLabel.Stack(), ptLabel.Stack(), pLabel.Stack()),
		container.NewGridWithColumns(2, bankHeading.Stack(), betHeading.Stack()),
		container.NewGridWithColumns(2, bankLabel.Stack(), betLabel.Stack()),
		widget.NewButton("Reset the Bank", reset),
		container.NewPadded(info),
	)

	return container.NewStack(container.NewVScroll(stuff), bg)
}

func setInfo(textBinding binding.String, pt *point.PointTracker) {
	if pt.NewPlayer == point.NEW_PLAYER {
		textBinding.Set("2, 3, 12 loses. 7, 11 wins, otherwise point is set")
	} else {
		textBinding.Set(fmt.Sprintf("7 loses, %d wins. Only Pass Line bets accepted.", pt.CurPoint.Value))
	}
}

func InfoText(textBinding binding.String, isDark bool) (info *canvas.Text) {
	textColor := color.Black
	if isDark {
		textColor = color.White
	}

	info = canvas.NewText("", textColor)
	info.TextSize = (3 * info.TextSize) / 4
	info.Alignment = fyne.TextAlignCenter
	return info
}

func updateDice(img []*canvas.Image, left int, images [][]image.Image, leftDie int, i int, right int, rightDie int) {
	img[left].Image = images[leftDie][i]
	img[left].Refresh()
	img[right].Image = images[rightDie][i]
	img[right].Refresh()
}

func handleKey(key string, bet BS, bank *Cash) {
	s, _ := bet.Get()
	if key == "DEL" {
		if len(s) > 0 {
			s = s[:len(s)-1]
			if canCover(s, bank) {
				bet.Set(s)
			}
		}
	} else if key == "AC" {
		bet.Set("")
	} else if key == "Bet 1/2" {
		setAuto(bet, bank.amt, 2)
	} else if key == "Bet 1/4" {
		setAuto(bet, bank.amt, 4)
	} else if key == "Bet All" {
		setAuto(bet, bank.amt, 1)
	} else if strings.ContainsAny(key[0:1], "0123456789.") {
		s += key
		if canCover(s, bank) {
			bet.Set(s)
		}
	}
}

func canCover(s string, bank *Cash) bool {
	amt := ToMoney(fromThousands(s))
	return amt <= bank.amt
}

type Money int64

func (m Money) String() string {
	return toThousands(strconv.FormatInt(int64(m), 10))
}

func ToMoney(s string) Money {
	i, _ := strconv.ParseInt(fromThousands(s), 10, 64)
	return Money(i)
}

type Cash struct {
	initialBank, amt Money
	amtString        BS
}

func NewCash(initialBank Money) *Cash {
	b := &Cash{
		initialBank: initialBank,
		amt:         initialBank,
		amtString:   NewBS(),
	}
	b.amtString.Set(b.String())
	return b
}

func (b *Cash) Reset(bet BS, initialBank Money) {
	b.initialBank = initialBank
	b.SetAmt(initialBank)
	half := initialBank / 2
	bet.Set(half.String())
}

func (b *Cash) AddBetAmt(bet BS, pos bool) {
	s, _ := bet.Get()
	betAmt := ToMoney(s)
	if !pos {
		betAmt = -betAmt
	}
	newAmt := b.amt + Money(betAmt)
	b.SetAmt(newAmt)
}

func (b *Cash) SetAmt(amt Money) {
	b.amt = amt
	fyne.CurrentApp().Preferences().SetInt("bank", int(amt))
	b.amtString.Set(b.String())
}

func (b *Cash) String() string {
	return b.amt.String()
}

func setAuto(bet BS, curAmt Money, factor Money) {
	if factor == 0 {
		factor = 1
	}
	betAmt := curAmt / factor
	bet.Set(betAmt.String())
}

func cacheImages(animationFiles embed.FS) [][]image.Image {
	images := make([][]image.Image, 6)
	for j := 1; j <= 6; j++ {
		images[j-1] = make([]image.Image, 0)
		for i := 70; i <= 140; i++ {
			fileName := fmt.Sprintf("media/Animation/roll-%d/%04d.png", j, i)
			data, err := animationFiles.ReadFile(fileName)
			if err != nil {
				panic(err)
			}
			img, err := png.Decode(bytes.NewReader(data))
			if err != nil {
				panic(err)
			}
			images[j-1] = append(images[j-1], img)
		}
	}
	return images
}

func toThousands(s string) string {
	var result []string
	for len(s) > 3 {
		result = append([]string{s[len(s)-3:]}, result...)
		s = s[:len(s)-3]
	}
	result = append([]string{s}, result...)
	return strings.Join(result, ",")
}

func fromThousands(s string) string {
	return strings.ReplaceAll(s, ",", "")
}

type BS struct{ binding.String }

func NewBS() BS { return BS{binding.NewString()} }

func (s BS) GetS() string { t, _ := s.Get(); return t }

type ThemedLabel struct {
	*widget.Label
	overlay *canvas.Rectangle
}

func NewThemedLabel(text string) *ThemedLabel {
	l := &ThemedLabel{Label: widget.NewLabel(text), overlay: canvas.NewRectangle(GREEN)}
	l.overlay.StrokeWidth = 1
	return l
}

func NewThemedLabelWithData(text BS) *ThemedLabel {
	l := &ThemedLabel{Label: widget.NewLabelWithData(text), overlay: canvas.NewRectangle(color.Transparent)}
	l.overlay.StrokeWidth = 1
	return l
}

func (t *ThemedLabel) Stack() fyne.CanvasObject {
	return container.NewStack(t.overlay, t.Label)
}
