// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	internalinterfaces "github.com/jenkins-x-labs/app-slack/pkg/client/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// SlackBots returns a SlackBotInformer.
	SlackBots() SlackBotInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// SlackBots returns a SlackBotInformer.
func (v *version) SlackBots() SlackBotInformer {
	return &slackBotInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
