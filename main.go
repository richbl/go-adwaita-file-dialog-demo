package main

import (
	"context"
	"log"
	"os"

	"github.com/diamondburned/gotk4-adwaita/pkg/adw"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

// App is the main application struct, using struct composition for the core adw.Application
type App struct {
	*adw.Application
	window *adw.ApplicationWindow
	status *gtk.Label
}

// Constants for the application
const (
	appID         = "gtk4-adwaita-file-open-demo"
	windowTitle   = "GTK4/Adwaita FileDialog Demo"
	defaultWidth  = 500
	defaultHeight = 300
)

// main is the application entry point
func main() {

	// Set up simple logging
	log.SetFlags(0)
	log.Println("Starting GTK4/Adwaita FileDialog demo...")

	// Create the application
	app := newApp()

	// Run the application, and then exit with the return code
	os.Exit(app.Run(os.Args))
}

// newApp creates a new application instance
func newApp() *App {

	// Create the application
	adwApp := adw.NewApplication(appID, gio.ApplicationFlagsNone)

	// Create the main application struct
	app := &App{
		Application: adwApp,
	}

	// Connect the Activate signal before running the application
	app.ConnectActivate(app.activate)

	return app
}

// activate is the application activation handler
func (a *App) activate() {

	// If the window already exists, just show it
	if a.window != nil {
		a.window.SetVisible(true)

		return
	}

	// Create the main application window
	a.window = adw.NewApplicationWindow(&a.Application.Application)
	a.window.SetTitle(windowTitle)
	a.window.SetDefaultSize(defaultWidth, defaultHeight)

	// Main structure: HeaderBar (top) + Content (bottom)
	rootBox := gtk.NewBox(gtk.OrientationVertical, 0)
	rootBox.Append(adw.NewHeaderBar())

	// Create the main content
	content := a.buildMainContent()
	content.SetVExpand(true)

	// Add the content to the root box
	rootBox.Append(content)

	// Set the root box as the content of the window
	a.window.SetContent(rootBox)
	a.window.SetVisible(true)

}

// buildMainContent creates the main content of the window
func (a *App) buildMainContent() *gtk.Box {

	// Create the main content box
	mainBox := gtk.NewBox(gtk.OrientationVertical, 12)
	mainBox.SetMarginTop(24)
	mainBox.SetMarginBottom(24)
	mainBox.SetMarginStart(24)
	mainBox.SetMarginEnd(24)

	// Title
	titleLabel := gtk.NewLabel("GTK4/Adwaita File Open Dialog Demo")
	titleLabel.AddCSSClass("title-1")
	mainBox.Append(titleLabel)

	// Description
	descLabel := gtk.NewLabel("This demonstrates the GTK4/Adwaita FileDialog.Open() method\nwith proper type casting and callback handling.")
	descLabel.AddCSSClass("body")
	descLabel.SetWrap(true)
	mainBox.Append(descLabel)

	// Open button
	openButton := gtk.NewButtonWithLabel("Open File")
	openButton.AddCSSClass("pill")
	openButton.AddCSSClass("suggested-action")
	openButton.SetIconName("document-open-symbolic")
	openButton.SetHExpand(false)
	openButton.ConnectClicked(a.handleOpenButton)
	mainBox.Append(openButton)

	// Status frame
	mainBox.Append(a.buildStatusFrame())

	return mainBox
}

// buildStatusFrame creates the status frame
func (a *App) buildStatusFrame() *gtk.Frame {

	// Create the status frame
	statusFrame := gtk.NewFrame("")
	statusFrame.AddCSSClass("card")
	statusFrame.SetMarginTop(24)

	// Create the status box
	statusBox := gtk.NewBox(gtk.OrientationVertical, 8)
	statusBox.SetMarginTop(12)
	statusBox.SetMarginBottom(12)
	statusBox.SetMarginStart(12)
	statusBox.SetMarginEnd(12)

	// Status label
	statusLabel := gtk.NewLabel("Selected File:")
	statusLabel.AddCSSClass("heading")
	statusBox.Append(statusLabel)

	// Status text
	a.status = gtk.NewLabel("No file selected")
	a.status.AddCSSClass("monospace")
	a.status.SetWrap(true)
	a.status.SetSelectable(true)
	a.status.SetXAlign(0)
	statusBox.Append(a.status)

	// Set the status box as the child of the status frame
	statusFrame.SetChild(statusBox)

	return statusFrame
}

// createFileFilters creates the ListStore of FileFilters
func createFileFilters() *gio.ListStore {

	// Text files filter
	textFilter := gtk.NewFileFilter()
	textFilter.SetName("Text files")
	textFilter.AddPattern("*.txt")
	textFilter.AddPattern("*.md")

	// Filters: Get the correct GType from an instance and create the ListStore
	filters := gio.NewListStore(textFilter.Type())
	filters.Append(textFilter.Object)

	// Set the files filter
	allFilter := gtk.NewFileFilter()
	allFilter.SetName("All files")
	allFilter.AddPattern("*")
	filters.Append(allFilter.Object)

	return filters
}

// handleOpenButton handles the Open button click
func (a *App) handleOpenButton() {

	log.Println("Opening file dialog...")

	// Create the file dialog
	dialog := gtk.NewFileDialog()
	dialog.SetTitle("Select a File")
	dialog.SetModal(true)

	// Use the helper to create filters
	dialog.SetFilters(createFileFilters())

	// Callback function for the Open button click
	dialog.Open(context.Background(), &a.window.Window, func(res gio.AsyncResulter) {

		// Open the file dialog
		file, err := dialog.OpenFinish(res)
		if err != nil {
			glib.IdleAdd(func() {
				a.showError("File Open Dialog Error", err.Error())
			})

			return
		}

		// Check for user interrupt
		if file == nil {
			glib.IdleAdd(func() {
				a.status.SetText("File Open Dialog cancelled by user")
			})

			return
		}

		// Success case: display the selected file
		path := file.Path()
		glib.IdleAdd(func() {
			a.status.SetText(path)
			log.Printf("Selected file: %s", path)
		})

	})

}

// showError shows an error message (using adw.AlertDialog)
func (a *App) showError(title, message string) {

	// Check if the window exists before presenting the dialog
	if a.window == nil {
		log.Printf("Cannot show error dialog (window is nil): %s: %s", title, message)

		return
	}

	// Show the error dialog
	dialog := adw.NewAlertDialog(title, message)
	dialog.AddResponse("ok", "OK")
	dialog.SetResponseAppearance("ok", adw.ResponseSuggested)
	dialog.SetCloseResponse("ok")
	dialog.Present(a.window)

}

// Run executes the application and returns the exit status code
func (a *App) Run(args []string) int {

	return a.Application.Run(args)
}
