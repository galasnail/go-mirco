package gomicro

import (
	"errors"
	"fmt"
)

type ServerPluginContainer struct {
	plugins []IPlugin
}

func (p *ServerPluginContainer)Add(plugin IPlugin) error {
	if p.plugins == nil {
		p.plugins = make([]IPlugin, 0)
	}
	pName := p.GetName(plugin)
	if pName != "" && p.GetByName(pName) != nil {
		return errors.New(fmt.Sprintf(("Cannot use the same plugin again, %s is already exists", pName))
	}

	p.plugins = append(p.plugins, plugin)
	return nil
}

func (p *ServerPluginContainer)GetName(plugin IPlugin) string {
	return plugin.Name()
}

// GetByName returns a plugin instance by it's name
func (p *ServerPluginContainer) GetByName(pluginName string) IPlugin {
	if p.plugins == nil {
		return nil
	}

	for _, plugin := range p.plugins {
		if plugin.Name() == pluginName {
			return plugin
		}
	}

	return nil
}

// Remove removes a plugin by it's name.
func (p *ServerPluginContainer) Remove(pluginName string) error {
	if p.plugins == nil {
		return errors.New(fmt.Sprintf("remove %s from plugins can not be nil", pluginName))
	}

	if pluginName == "" {
		//return error: cannot delete an unamed plugin
		return errors.New("remove plugins")
	}

	indexToRemove := -1
	for i := range p.plugins {
		if p.GetName(p.plugins[i]) == pluginName {
			indexToRemove = i
			break
		}
	}
	if indexToRemove == -1 {
		return ErrPluginRemoveNotFound.Return()
	}

	p.plugins = append(p.plugins[:indexToRemove], p.plugins[indexToRemove+1:]...)

	return nil
}


// GetAll returns all activated plugins
func (p *ServerPluginContainer) GetAll() []IPlugin {
	return p.plugins
}

// DoRegister invokes DoRegister plugin.
func (p *ServerPluginContainer) DoRegister(name string, rcvr interface{}, metadata ...string) error {
	var es []error
	for i := range p.plugins {

		if plugin, ok := p.plugins[i].(IRegisterPlugin); ok {
			err := plugin.Register(name, rcvr, metadata...)
			if err != nil {
				errors = append(errors, err)
			}
		}
	}

	if len(es) > 0 {
		return errors.New("multi errors!")
	}
	return nil
}


type (
	IServerPluginContainer interface {
		Add(plugin IPlugin) error
		Remove(pluginName string) error
		GetName(plugin IPlugin) string
		GetByName(pluginName string) IPlugin
		GetAll() []IPlugin

		DoRegister(name string, rcvr interface{}, metadata ...string) error

		/*
			DoPostConnAccept(conn net.Conn) (net.Conn, bool)

			DoPreReadRequestHeader(r *grpc.Stream) error
			DoPostReadRequestHeader(r *grpc.Request) error

			DoPreReadRequestBody(body interface{}) error
			DoPostReadRequestBody(body interface{}) error

			DoPreWriteResponse(*grpc.Response, interface{}) error
			DoPostWriteResponse(*grpc.Response, interface{}) error
		*/
	}

	//IRegisterPlugin represents register plugin.
	IRegisterPlugin interface {
		Register(name string, rcvr interface{}, metadata ...string) error
	}
)






















