package controllers

import (
	"context"

	"github.com/pkg/errors"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func getPullSecret(ctx context.Context, c client.Client, name, namespace string) (string, error) {
	secret := &corev1.Secret{}
	key := types.NamespacedName{
		Namespace: namespace,
		Name:      name,
	}
	if err := c.Get(ctx, key, secret); err != nil {
		return "", errors.Wrapf(err, "failed to get pull secret %s", key)
	}

	data, ok := secret.Data["pullSecret"]
	if !ok {
		return "", errors.Errorf("pullSecret is not configured")
	}

	return string(data), nil
}
