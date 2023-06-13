package router

import (
	"cloud.google.com/go/storage"
	"github.com/slack-go/slack"

	"k8s.io/test-infra/prow/config"

	"k8s.io/client-go/kubernetes"

	"github.com/openshift/ci-tools/pkg/slack/events"
	"github.com/openshift/ci-tools/pkg/slack/events/helpdesk"
	"github.com/openshift/ci-tools/pkg/slack/events/joblink"
	"github.com/openshift/ci-tools/pkg/slack/events/mention"
)

// ForEvents returns a Handler that appropriately routes
// event callbacks for the handlers we know about
func ForEvents(client *slack.Client, kubeClient kubernetes.Interface, config config.Getter, gcsClient *storage.Client, keywordsConfig helpdesk.KeywordsConfig, helpdeskAlias, forumChannelId string, requireWorkflowsInForum bool) events.Handler {
	return events.MultiHandler(
		helpdesk.MessageHandler(client, keywordsConfig, helpdeskAlias, forumChannelId, requireWorkflowsInForum),
		helpdesk.FAQHandler(client, kubeClient, forumChannelId),
		mention.Handler(client),
		joblink.Handler(client, joblink.NewJobGetter(config), gcsClient),
	)
}
