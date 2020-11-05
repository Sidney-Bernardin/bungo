package bungo

func basedOnComponent(params map[string]string) interface{} {

	var component string

	if _, ok := params["components"]; !ok {
		component = "0"
	}

	component = params["components"]

	switch component {

	case "0":
		return struct{}{}

	case "200":
		return SingleComponentResponseOfDestinyCharacterComponent{}

	case "205":
		return SingleComponentResponseOfDestinyInventoryComponent{}
	}

	return nil
}
