package exp

import (
	"bytes"
	"craps/opts"
	"embed"
	"fmt"
	"image"
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

func App3(animationFiles embed.FS, opt opts.Options) *fyne.Container {
	images := cacheImages(animationFiles)

	resultString := binding.NewString()
	resultString.Set("Please roll.")
	result := widget.NewLabelWithData(resultString)

	// widgets showing die animation side-by-side
	img := make([]*canvas.Image, 2)

	left, right := 0, 1
	for i := range []int{left, right} {
		img[i] = canvas.NewImageFromImage(images[i][len(images)-1])
		img[i].FillMode = canvas.ImageFillOriginal
		img[i].ScaleMode = canvas.ImageScaleFastest
	}

	initialBank := float32(100.0)
	bank := NewCash(100.0)
	bankLabel := widget.NewLabelWithData(bank.amtString)
	bet := binding.NewString()
	bet.Set(strconv.FormatFloat(float64(initialBank)/2, 'f', 2, 32))
	betLabel := widget.NewLabelWithData(bet)

	// the zero based value of the rolled die
	leftDie, rightDie := 0, 0

	var rollButton *widget.Button
	doAnimation := func(tick float32) {
		// there are len(images) to display in 4s, tick will be 0.5 at 2s for instance, which is len(images)/2, so the image # is tick*len(images)
		i := int(tick * float32(len(images[0])-1))
		updateDice(img, left, images, leftDie, i, right, rightDie)
		if tick == 1.0 {
			resultString.Set("Please roll.")
			rollButton.Enable()
			total := leftDie + 1 + rightDie + 1 // Dice values are 1-indexed
			resultText := fmt.Sprintf("You rolled: %d", total)
			switch total {
			case 7, 11:
				resultText += ". You Win!"
				bank.AddBetAmt(bet, false)
			case 2, 3, 12:
				resultText += ". You Lose."
				bank.AddBetAmt(bet, true)
				if bank.amt <= 0.0 {
					bank.SetAmt(bank.initialBank)
					bet.Set("0.00")
					resultText += " Refreshed bank."
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
		f, _ := strconv.ParseFloat(b, 32)
		bet.Set(strconv.FormatFloat(f, 'f', 2, 32))
		leftDie, rightDie = rand.Intn(6), rand.Intn(6)
		resultString.Set("Rolling...")
		fyne.NewAnimation(1*time.Second, doAnimation).Start()
	})

	keys := make([]fyne.CanvasObject, 0)
	for _, key := range []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", ".", "0", "/", "AC", "Auto", "DEL"} {
		b := widget.NewButton(" "+key+" ", func() { handleKey(key, bet, bank) })
		keys = append(keys, b)
	}

	return container.NewVBox(
		container.NewHBox(
			layout.NewSpacer(),
			img[left],
			img[right],
			layout.NewSpacer(),
		),
		container.NewGridWithColumns(3, keys...),
		rollButton,
		result,
		bankLabel,
		betLabel,
	)
}

func updateDice(img []*canvas.Image, left int, images [][]image.Image, leftDie int, i int, right int, rightDie int) {
	img[left].Image = images[leftDie][i]
	img[left].Refresh()
	img[right].Image = images[rightDie][i]
	img[right].Refresh()
}

func handleKey(key string, bet binding.String, bank *Cash) {
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
	amt, err := strconv.ParseFloat(s, 32)
	return err == nil && amt <= float64(bank.amt)
}

type Cash struct {
	initialBank, amt float32
	amtString        binding.String
}

func NewCash(initialBank float32) *Cash {
	b := &Cash{
		initialBank: initialBank,
		amt:         initialBank,
		amtString:   binding.NewString(),
	}
	b.amtString.Set(b.String())
	return b
}

func (b *Cash) AddBetAmt(bet binding.String, neg bool) {
	s, _ := bet.Get()
	betAmt, _ := strconv.ParseFloat(s, 32)
	if neg {
		betAmt = -betAmt
	}
	newAmt := b.amt + float32(betAmt)
	b.SetAmt(newAmt)
}

func (b *Cash) SetAmt(amt float32) {
	b.amt = amt
	b.amtString.Set(b.String())
}

func (b *Cash) String() string {
	return fmt.Sprintf("%.2f", b.amt)
}

func setAuto(bet binding.String, curAmt float32) {
	betAmt := curAmt / 2
	bet.Set(strconv.FormatFloat(float64(betAmt), 'f', 2, 32))
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
