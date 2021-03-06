// Code generated by entc, DO NOT EDIT.

package ent

import "context"

func (u *User) UserLogs(ctx context.Context) ([]*UserLog, error) {
	result, err := u.Edges.UserLogsOrErr()
	if IsNotLoaded(err) {
		result, err = u.QueryUserLogs().All(ctx)
	}
	return result, err
}

func (u *User) UserOaInfo(ctx context.Context) ([]*UserOaInfo, error) {
	result, err := u.Edges.UserOaInfoOrErr()
	if IsNotLoaded(err) {
		result, err = u.QueryUserOaInfo().All(ctx)
	}
	return result, err
}

func (u *User) UserOpInfo(ctx context.Context) ([]*UserOpInfo, error) {
	result, err := u.Edges.UserOpInfoOrErr()
	if IsNotLoaded(err) {
		result, err = u.QueryUserOpInfo().All(ctx)
	}
	return result, err
}
