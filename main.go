// AntiICloud — это шуточная утилита.
// Она НЕ выполняет реальный сброс iCloud, НЕ отправляет данные в сеть
// и НЕ совершает никаких вредоносных действий.
// Программа просто рисует в консоли правдоподобный процесс и в конце печатает "Test".

package main

import (
	"bufio"
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"strings"
	"time"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	done := make(chan struct{})
	go func() {
		runPrank(ctx)
		close(done)
	}()

	select {
	case <-done:
	case <-ctx.Done():
	}

	fmt.Println()
	fmt.Println("Test")
	fmt.Println()
	time.Sleep(300 * time.Millisecond)
}

func runPrank(ctx context.Context) {
	fmt.Println()
	fmt.Println("AntiICloud v2.0")
	fmt.Println("Device Lock Management Utility")
	fmt.Println(strings.Repeat("=", 50))
	fmt.Println()
	fmt.Println("[!] DISCLAIMER: use only on devices you own or have")
	fmt.Println("    explicit permission to manage. Any other use is prohibited.")
	fmt.Println()

	_, ok := readLine(ctx, "[?] Enter device serial / IMEI: ")
	if !ok {
		return
	}

	_, ok = readEnter(ctx, "[?] Connect the device via USB and press ENTER...")
	if !ok {
		return
	}

	if !spinner(ctx, "Authenticating with Apple GSX API", 2*time.Second) {
		return
	}
	if !spinner(ctx, "Fetching activation lock status", 2*time.Second) {
		return
	}
	fmt.Println("[*] Activation lock status: ACTIVE")
	if !spinner(ctx, "Loading bypass payload", 2*time.Second) {
		return
	}
	if !progressBar(ctx, "Bypassing activation lock", 15*time.Second) {
		return
	}
	fmt.Println("[*] Bypass complete. Do not disconnect the device.")
	fmt.Println("[*] Waiting for the device to reboot and finalize...")
	fmt.Println()

	countdown(ctx, 60)
}

func readLine(ctx context.Context, prompt string) (string, bool) {
	fmt.Print(prompt)
	ch := make(chan string, 1)
	go func() {
		var s string
		_, _ = fmt.Scanln(&s)
		ch <- s
	}()
	select {
	case s := <-ch:
		return s, true
	case <-ctx.Done():
		return "", false
	}
}

func readEnter(ctx context.Context, prompt string) (bool, bool) {
	fmt.Print(prompt)
	ch := make(chan bool, 1)
	go func() {
		_, _ = bufio.NewReader(os.Stdin).ReadString('\n')
		ch <- true
	}()
	select {
	case <-ch:
		return true, true
	case <-ctx.Done():
		return false, false
	}
}

func spinner(ctx context.Context, label string, duration time.Duration) bool {
	frames := `|/-\`
	start := time.Now()
	i := 0
	for time.Since(start) < duration {
		select {
		case <-ctx.Done():
			return false
		default:
		}
		fmt.Printf("\r[*] %s %c", label, frames[i%len(frames)])
		time.Sleep(80 * time.Millisecond)
		i++
	}
	fmt.Printf("\r[*] %s ... OK\n", label)
	return true
}

func progressBar(ctx context.Context, label string, total time.Duration) bool {
	steps := 40
	stepDuration := total / time.Duration(steps)
	fmt.Printf("[*] %s [", label)
	for i := 0; i <= steps; i++ {
		select {
		case <-ctx.Done():
			return false
		case <-time.After(stepDuration + time.Duration(rand.Intn(150))*time.Millisecond):
		}
		pct := i * 100 / steps
		filled := strings.Repeat("#", i)
		empty := strings.Repeat("-", steps-i)
		fmt.Printf("\r[*] %s [%s%s] %d%%", label, filled, empty, pct)
	}
	fmt.Println()
	return true
}

func countdown(ctx context.Context, seconds int) bool {
	for i := seconds; i >= 0; i-- {
		select {
		case <-ctx.Done():
			return false
		case <-time.After(1 * time.Second):
		}
		fmt.Printf("\r[*] Finalizing in %2ds...", i)
	}
	fmt.Println()
	return true
}
