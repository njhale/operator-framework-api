/*
Copyright 2019 Red Hat, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
// Code generated by client-gen. DO NOT EDIT.

package v2alpha1

import (
	"github.com/operator-framework/api/generated/clientset/versioned/scheme"
	v2alpha1 "github.com/operator-framework/api/operators/v2alpha1"
	rest "k8s.io/client-go/rest"
)

type OperatorsV2alpha1Interface interface {
	RESTClient() rest.Interface
	OperatorsGetter
}

// OperatorsV2alpha1Client is used to interact with features provided by the operators.coreos.com group.
type OperatorsV2alpha1Client struct {
	restClient rest.Interface
}

func (c *OperatorsV2alpha1Client) Operators() OperatorInterface {
	return newOperators(c)
}

// NewForConfig creates a new OperatorsV2alpha1Client for the given config.
func NewForConfig(c *rest.Config) (*OperatorsV2alpha1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &OperatorsV2alpha1Client{client}, nil
}

// NewForConfigOrDie creates a new OperatorsV2alpha1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *OperatorsV2alpha1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new OperatorsV2alpha1Client for the given RESTClient.
func New(c rest.Interface) *OperatorsV2alpha1Client {
	return &OperatorsV2alpha1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v2alpha1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *OperatorsV2alpha1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
