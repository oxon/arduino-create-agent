//
//  trayicon.go
//
//  Created by Martino Facchin
//  Copyright (c) 2015 Arduino LLC
//
//  Permission is hereby granted, free of charge, to any person
//  obtaining a copy of this software and associated documentation
//  files (the "Software"), to deal in the Software without
//  restriction, including without limitation the rights to use,
//  copy, modify, merge, publish, distribute, sublicense, and/or sell
//  copies of the Software, and to permit persons to whom the
//  Software is furnished to do so, subject to the following
//  conditions:
//
//  The above copyright notice and this permission notice shall be
//  included in all copies or substantial portions of the Software.
//
//  THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
//  EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
//  OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
//  NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT
//  HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,
//  WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
//  FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
//  OTHER DEALINGS IN THE SOFTWARE.
//

// +build !cli

package main

import (
	"runtime"

	log "github.com/Sirupsen/logrus"
	"github.com/oxon/arduino-create-agent/icon"
	"github.com/getlantern/systray"
	"github.com/skratchdot/open-golang/open"
)

func setupSysTray() {
	runtime.LockOSThread()
	if *hibernate == true {
		systray.Run(setupSysTrayHibernate)
	} else {
		systray.Run(setupSysTrayReal)
	}
}

func addRebootTrayElement() {
	reboot_tray := systray.AddMenuItem("Reboot to update", "")

	go func() {
		<-reboot_tray.ClickedCh
		systray.Quit()
		log.Println("Restarting now...")
		log.Println("Restart because addReebotTrayElement")
		restart("")
	}()
}

func setupSysTrayReal() {

	systray.SetIcon(icon.GetIcon())
	mProgram := systray.AddMenuItem("Program Oxocard", "OxocardAgent")
	mUrl := systray.AddMenuItem("Go to Oxocard-Website", "OxocardAgent")
	menuVer := systray.AddMenuItem("OxocardAgent version "+version, "")
	mQuit := systray.AddMenuItem("Quit Agent", "")

	menuVer.Disable()

	go func() {
		<-mQuit.ClickedCh
		systray.Quit()
		exit()
	}()

	go func() {
		for {
			<-mPorgram.ClickedCh
			open.Start("http://oxocard.ch/oxocard-blockly")
		}
	}()

	go func() {
		for {
			<-mUrl.ClickedCh
			open.Start("http://oxocard.ch")
		}
	}()
}

func setupSysTrayHibernate() {

	systray.SetIcon(icon.GetIconHiber())
	mOpen := systray.AddMenuItem("Open Plugin", "")
	mQuit := systray.AddMenuItem("Kill Plugin", "")

	go func() {
		<-mOpen.ClickedCh
		*hibernate = false
		log.Println("Restart for hubernation")
		systray.Quit()
		restart("")
	}()

	go func() {
		<-mQuit.ClickedCh
		systray.Quit()
		exit()
	}()
}

func quitSysTray() {
	systray.Quit()
}
