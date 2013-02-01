package main

import "github.com/guelfey/go.dbus"

func main() {
	conn, err := dbus.ConnectSessionBus()
	if err != nil {
		panic(err)
	}
	obj := conn.Object("org.freedesktop.Notifications", "/org/freedesktop/Notifications")
	reply := <-obj.Call("org.freedesktop.Notifications.Notify", 0, "", uint32(0),
		"", "Test", "This is a test of the DBus bindings for go.", []string{},
			map[string]dbus.Variant{}, int32(5000))
	if reply.Err != nil {
		panic(reply.Err)
	}
}