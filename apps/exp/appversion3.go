package exp

import (
	"bytes"
	"craps/opts"
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
	"fyne.io/fyne/v2/widget"
	"golang.org/x/exp/rand"
)

var RED, GREEN = color.RGBA{255, 0, 0, 128}, color.RGBA{0, 255, 0, 128}

func App3(animationFiles embed.FS, opt opts.Options) *fyne.Container {
	initialMsg := "Welcome to Simple Craps."
	images := cacheImages(animationFiles)

	resultString := NewBS()
	resultString.Set(initialMsg)
	result := NewThemedLabelWithData(resultString)

	// widgets showing die animation side-by-side
	img := make([]*canvas.Image, 2)

	left, right := 0, 1
	for i := range []int{left, right} {
		img[i] = canvas.NewImageFromImage(images[i][len(images)-1])
		img[i].FillMode, img[i].ScaleMode = canvas.ImageFillOriginal, canvas.ImageScaleFastest
	}

	initialBank := Money(2000)
	bank := NewCash(initialBank)
	bankLabel := NewThemedLabelWithData(bank.amtString)
	bankLabel.overlay.StrokeColor, bankLabel.overlay.StrokeWidth = GREEN, 2

	bet := NewBS()
	bet.Set(Money(int64(initialBank) / 2).String())
	betLabel := NewThemedLabelWithData(bet)
	betLabel.overlay.StrokeColor, betLabel.overlay.StrokeWidth = GREEN, 2

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
			switch total {
			case 7, 11:
				resultText += ". You Win!"
				bank.AddBetAmt(bet, true)
			case 2, 3, 12:
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
				resultText += ". Push."
			}
			resultString.Set(resultText)
		}
	}

	rand.Seed(uint64(time.Now().UnixNano()))
	rollButton = widget.NewButton("Roll", func() {
		rollButton.Disable()
		b, _ := bet.Get()
		i := ToMoney(b)
		bet.Set(i.String())
		leftDie, rightDie = rand.Intn(6), rand.Intn(6)
		resultString.Set("Rolling...")
		fyne.NewAnimation(1*time.Second, doAnimation).Start()
	})

	keys := make([]fyne.CanvasObject, 0)
	for _, key := range []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", " ", "0", " ", "AC", "Auto", "DEL"} {
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

	stuff := container.NewVBox(
		dice,
		container.NewGridWithColumns(3, keys...),
		rollButton,
		container.NewGridWithColumns(2, bankHeading.Stack(), betHeading.Stack()),
		container.NewGridWithColumns(2, bankLabel.Stack(), betLabel.Stack()),
		result.Stack(),
	)

	return container.NewStack(container.NewVScroll(stuff), bg)
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
	} else if strings.ContainsAny(key, "0123456789.") {
		s += key
		if canCover(s, bank) {
			bet.Set(s)
		}
	} else if key == "Auto" {
		setAuto(bet, bank.amt)
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
	b.amtString.Set(b.String())
}

func (b *Cash) String() string {
	return b.amt.String()
}

func setAuto(bet BS, curAmt Money) {
	betAmt := curAmt / 2
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
