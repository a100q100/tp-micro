// Code generated by 'micro gen' command.
// DO NOT EDIT!

package api

import (
	tp "github.com/henrylee2cn/teleport"
)

// Route registers handlers to router.
func Route(_root string, _router *tp.Router) {
	// root router group
	_group := _router.SubRoute(_root)

	// custom router
	customRoute(_group.ToRouter())

	// automatically generated router

	// PULL APIs...
	{
		_group.RouteCallFunc(Home)
		_group.RouteCall(new(Math))
	}

	// PUSH APIs...
	{
		_group.RoutePushFunc(Stat)
	}
}
