package reqflags

import (
	"gopkg.in/urfave/cli.v2"
)

// OnMissing is 欠けているフラグ名とリカバリ用関数のマップ.
type OnMissing map[string]func() error

// Recover is フラグが欠けていたらそのフラグを入力させたい.
func Recover(ctx *cli.Context, onMissing OnMissing) error {
	for maybeMissing, onMissingFunc := range onMissing {
		if ctx.IsSet(maybeMissing) {
			continue
		}
		if err := onMissingFunc(); err != nil {
			return err
		}
	}
	return nil
}

// IsSufficient is フラグが全てセットされていればtrue, nil. そうでなければfalse, 欠けているフラグの名前が返る.
func IsSufficient(ctx *cli.Context, flags []string) (bool, []string) {
	missingFlags := make([]string, 0, len(flags))
	for _, flagName := range flags {
		if !ctx.IsSet(flagName) {
			missingFlags = append(missingFlags, flagName)
		}
	}

	if len(missingFlags) > 0 {
		return false, missingFlags
	}
	return true, nil
}
