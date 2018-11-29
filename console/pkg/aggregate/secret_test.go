package aggregate

import (
	"github.com/magiconair/properties/assert"
	"hidevops.io/hioak/starter/kube"
	"hidevops.io/mio/console/pkg/command"
	"k8s.io/client-go/kubernetes/fake"
	"testing"
)

func TestSecretCreate(t *testing.T) {
	client := fake.NewSimpleClientset()
	s := kube.NewSecret(client)
	ra := NewSecretAggregate(s)
	secret := &command.Secret{
		Name:      "hello-world",
		Namespace: "demo",
	}
	err := ra.Create(secret)
	assert.Equal(t, nil, err)
	_, err = ra.Get("hello-world", "demo")
	assert.Equal(t, nil, err)
}
