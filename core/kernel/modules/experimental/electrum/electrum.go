package electrum

import (
	"fmt"
	"vnh1/types"

	v8 "rogchap.com/v8go"
)

type ElectrumModule struct {
	context *v8.Context
}

func (o *ElectrumModule) Init(kernel types.KernelInterface, iso *v8.Isolate, context *v8.Context) error {
	// Das Consolen Objekt wird definiert
	console := v8.NewObjectTemplate(iso)

	// Der Kontext wird abgespeichert
	o.context = v8.NewContext(iso)

	// Das Console Objekt wird final erzeugt
	consoleObj, err := console.NewInstance(o.context)
	if err != nil {
		return fmt.Errorf("Kernel->_new_kernel_load_console_module: " + err.Error())
	}

	// Das Objekt wird als Import Registriert
	if err := kernel.AddImportModule("net", consoleObj.Value); err != nil {
		return fmt.Errorf("ModuleNetwork->Init: " + err.Error())
	}

	// Kein Fehler
	return nil
}

func (o *ElectrumModule) OnlyForMain() bool {
	return false
}

func (o *ElectrumModule) GetName() string {
	return "console"
}

func NewElectrumModule() *ElectrumModule {
	return new(ElectrumModule)
}
