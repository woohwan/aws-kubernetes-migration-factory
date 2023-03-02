package resource

import (
	app "k8s.io/api/apps/v1"
	autoscaling "k8s.io/api/autoscaling/v1"
	v1 "k8s.io/api/core/v1"
	networking "k8s.io/api/networking/v1"
	podsecuritypolicy "k8s.io/api/policy/v1beta1"
	rbac "k8s.io/api/rbac/v1"
	storage "k8s.io/api/storage/v1"

	admissionregistration "k8s.io/api/admissionregistration/v1"
	batchv1 "k8s.io/api/batch/v1"
	batchv1beta1 "k8s.io/api/batch/v1beta1"
)
type Resources struct {
	Svcl				[]v1.Service
	Nsl					*v1.NamespaceList
	Dsl					[]app.DaemonSet
	SecretList 	[]v1.Secret

	// var map[string]
	Depl																[]app.Deployment
	StorageClassList										[]storage.StorageClass
	ConfigMapsList												[]v1.ConfigMap
	IngressList													[]networking.Ingress
	RoleList														[]rbac.Role
	RoleBindingList											[]rbac.RoleBinding
	ClusterRoleList											[]rbac.ClusterRole
	ClusterRoleBindingList							[]rbac.ClusterRoleBinding
	HpaList															[]autoscaling.HorizontalPodAutoscaler
	PspList															[]podsecuritypolicy.PodSecurityPolicy
	SvcAccList													[]v1.ServiceAccount
	CronJobList													[]batchv1beta1.CronJob
	JobList															[]batchv1.Job
	PersistentVolumeClaimsList						[]v1.PersistentVolumeClaim
	MutatingWebhookConfigurationList 		[]admissionregistration.MutatingWebhookConfiguration
	ValidatingWebhookConfigurationList	[]admissionregistration.ValidatingWebhookConfiguration
	HelmList														map[string]map[string]string
}