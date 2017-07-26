package kingpin

// Action callback executed at various stages after all values are populated.
// The application, commands, arguments and flags all have corresponding
// actions.
type Action func(*ParseContext) error

type actionMixin struct {
	actions     []Action
	preActions  []Action
	postActions []Action
}

type actionApplier interface {
	applyActions(*ParseContext) error
	applyPreActions(*ParseContext) error
}

func (a *actionMixin) addAction(action Action) {
	a.actions = append(a.actions, action)
}

func (a *actionMixin) addPreAction(action Action) {
	a.preActions = append(a.preActions, action)
}

func (a *actionMixin) addPostAction(action Action) {
	a.postActions = append(a.postActions, action)
}

func (a *actionMixin) applyActions(context *ParseContext) error {
	for _, action := range a.actions {
		if err := action(context); err != nil {
			return err
		}
	}
	return nil
}

func (a *actionMixin) applyPreActions(context *ParseContext) error {
	for _, preAction := range a.preActions {
		if err := preAction(context); err != nil {
			return err
		}
	}
	return nil
}

func (a *actionMixin) applyPostActions(context *ParseContext) error {
	for _, postActions := range a.postActions {
		if err := postActions(context); err != nil {
			return err
		}
	}
	return nil
}
