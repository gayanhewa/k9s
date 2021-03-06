package views

import (
	"fmt"

	"github.com/derailed/k9s/internal/resource"
	"github.com/gdamore/tcell"
)

type cronJobView struct {
	*resourceView
}

func newCronJobView(t string, app *appView, list resource.List, c colorerFn) resourceViewer {
	v := cronJobView{
		resourceView: newResourceView(t, app, list, c).(*resourceView),
	}
	v.extraActionsFn = v.extraActions
	v.switchPage("cronjob")
	return &v
}

func (v *cronJobView) trigger(evt *tcell.EventKey) *tcell.EventKey {
	if !v.rowSelected() {
		return evt
	}

	v.app.flash(flashInfo, fmt.Sprintf("Triggering %s %s", v.list.GetName(), v.selectedItem))
	if err := v.list.Resource().(resource.Runner).Run(v.selectedItem); err != nil {
		v.app.flash(flashErr, "Boom!", err.Error())
		return evt
	}
	return nil
}

func (v *cronJobView) extraActions(aa keyActions) {
	aa[tcell.KeyCtrlT] = newKeyAction("Trigger", v.trigger, true)
}
