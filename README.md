# Docustay for Go

The official Go client for the [Docustay](https://docustay.app) e-signature API: send a document from a template, get a webhook when each person signs, download the sealed PDF.

- API version: 2026-10-07
- Package version: 0.2.2

## Install

```sh
go get github.com/docustay/docustay-go/docustay@v0.2.2
```

The module needs Go 1.23 or newer.

## Send your first document

Put your key in `DOCUSTAY_API_KEY` and a template id in `DOCUSTAY_TEMPLATE_ID`.

```go
cfg := docustay.NewConfiguration()
client := docustay.NewAPIClient(cfg)
ctx := context.WithValue(context.Background(), docustay.ContextAccessToken, os.Getenv("DOCUSTAY_API_KEY"))

request := *docustay.NewSendRequest(os.Getenv("DOCUSTAY_TEMPLATE_ID"), []docustay.SendRequestPartiesInner{
	*docustay.NewSendRequestPartiesInner("signer_1", "Pat Example", "pat@example.com"),
})
doc, _, err := client.DocumentsAPI.SendDocument(ctx).SendRequest(request).IdempotencyKey("send-0001").Execute()
```

More: https://docustay.app/docs

## API reference

All URIs are relative to *https://app.docustay.app*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AccountAPI* | [**AiUsage**](docs/AccountAPI.md#aiusage) | **Get** /api/v1/ai/usage | AI drafts used against your plan
*AccountAPI* | [**PaymentsStatus**](docs/AccountAPI.md#paymentsstatus) | **Get** /api/v1/payments/status | Is Stripe ready to take a payment?
*AccountAPI* | [**Preflight**](docs/AccountAPI.md#preflight) | **Post** /api/v1/preflight | Ask what an action would do to your plan, before doing it
*AccountAPI* | [**Usage**](docs/AccountAPI.md#usage) | **Get** /api/v1/usage | Documents, templates and emails used against your plan
*DocumentsAPI* | [**AuditLogCsv**](docs/DocumentsAPI.md#auditlogcsv) | **Get** /api/v1/documents/{id}/audit-log.csv | The document&#39;s audit log as CSV
*DocumentsAPI* | [**CancelAiJob**](docs/DocumentsAPI.md#cancelaijob) | **Post** /api/v1/documents/ai/jobs/{id}/cancel | Cancel an AI job
*DocumentsAPI* | [**CheckDocument**](docs/DocumentsAPI.md#checkdocument) | **Post** /api/v1/documents/check | Check a document&#39;s words before making it
*DocumentsAPI* | [**CreateDocument**](docs/DocumentsAPI.md#createdocument) | **Post** /api/v1/documents | Make a document from words and people
*DocumentsAPI* | [**DeleteDraft**](docs/DocumentsAPI.md#deletedraft) | **Delete** /api/v1/documents/{id} | Delete a draft that was never sent
*DocumentsAPI* | [**DetectFields**](docs/DocumentsAPI.md#detectfields) | **Post** /api/v1/documents/detect-fields | Find the fields on a PDF
*DocumentsAPI* | [**DownloadSignedPdf**](docs/DocumentsAPI.md#downloadsignedpdf) | **Get** /api/v1/documents/{id}/signed.pdf | Download the signed copy
*DocumentsAPI* | [**DraftWithAi**](docs/DocumentsAPI.md#draftwithai) | **Post** /api/v1/documents/ai/draft | Ask for an AI-written draft
*DocumentsAPI* | [**EditWithAi**](docs/DocumentsAPI.md#editwithai) | **Post** /api/v1/documents/ai/edit | Ask the assistant to change a document
*DocumentsAPI* | [**EmbedSession**](docs/DocumentsAPI.md#embedsession) | **Post** /api/v1/documents/{id}/embed-session | Start an embedded signing session for a seat
*DocumentsAPI* | [**GetAiDraft**](docs/DocumentsAPI.md#getaidraft) | **Get** /api/v1/documents/ai/jobs/{id} | Read an AI draft job
*DocumentsAPI* | [**GetDocument**](docs/DocumentsAPI.md#getdocument) | **Get** /api/v1/documents/{id} | One document and where each person is
*DocumentsAPI* | [**GetDraftSource**](docs/DocumentsAPI.md#getdraftsource) | **Get** /api/v1/documents/{id}/source | The words, people and choices of a draft
*DocumentsAPI* | [**GetWizard**](docs/DocumentsAPI.md#getwizard) | **Get** /api/v1/ai/wizard | The wizard&#39;s questions
*DocumentsAPI* | [**ListDocuments**](docs/DocumentsAPI.md#listdocuments) | **Get** /api/v1/documents | List documents
*DocumentsAPI* | [**ListTemplates**](docs/DocumentsAPI.md#listtemplates) | **Get** /api/v1/documents/templates | Templates you can send
*DocumentsAPI* | [**Remind**](docs/DocumentsAPI.md#remind) | **Post** /api/v1/documents/{id}/remind | Remind the people who still have to sign
*DocumentsAPI* | [**ReviseDocument**](docs/DocumentsAPI.md#revisedocument) | **Post** /api/v1/documents/{id}/revise | Change the words of a draft
*DocumentsAPI* | [**Search**](docs/DocumentsAPI.md#search) | **Get** /api/v1/search | Search documents, templates and people
*DocumentsAPI* | [**SendDocument**](docs/DocumentsAPI.md#senddocument) | **Post** /api/v1/documents/send | Send a template to people to sign
*DocumentsAPI* | [**SendDraft**](docs/DocumentsAPI.md#senddraft) | **Post** /api/v1/documents/{id}/send | Send a draft
*DocumentsAPI* | [**TestSign**](docs/DocumentsAPI.md#testsign) | **Post** /api/v1/documents/{id}/test-sign | Sign a test document (test mode only)
*DocumentsAPI* | [**VoidDocument**](docs/DocumentsAPI.md#voiddocument) | **Post** /api/v1/documents/{id}/void | Void a document that is still out for signature
*EmailAPI* | [**ListEmails**](docs/EmailAPI.md#listemails) | **Get** /api/v1/emails | The log of email about your documents
*EmbeddingAPI* | [**BuilderSession**](docs/EmbeddingAPI.md#buildersession) | **Post** /api/v1/embedded/builder-session | Start an embedded template builder session
*EmbeddingAPI* | [**ScribeSession**](docs/EmbeddingAPI.md#scribesession) | **Post** /api/v1/embedded/scribe-session | Start an embedded Scribe session
*PaymentsAPI* | [**ListPayments**](docs/PaymentsAPI.md#listpayments) | **Get** /api/v1/payments | List payments taken through documents
*PaymentsAPI* | [**RefundPayment**](docs/PaymentsAPI.md#refundpayment) | **Post** /api/v1/payments/{id}/refund | Refund a payment, in full or in part
*PlaysAPI* | [**AnswerDraft**](docs/PlaysAPI.md#answerdraft) | **Post** /api/v1/drafts/{draftRef}/answers | Answer a template&#39;s questions
*PlaysAPI* | [**ArchivePlay**](docs/PlaysAPI.md#archiveplay) | **Delete** /api/v1/plays/{name} | Archive a play
*PlaysAPI* | [**DuplicatePlay**](docs/PlaysAPI.md#duplicateplay) | **Post** /api/v1/plays/{name}/duplicate | Copy a play under a new name
*PlaysAPI* | [**GetPlay**](docs/PlaysAPI.md#getplay) | **Get** /api/v1/plays/{name} | One play
*PlaysAPI* | [**ListPlaybooks**](docs/PlaysAPI.md#listplaybooks) | **Get** /api/v1/playbooks | List playbooks (packs)
*PlaysAPI* | [**ListPlays**](docs/PlaysAPI.md#listplays) | **Get** /api/v1/plays | List plays
*PlaysAPI* | [**PublishPlay**](docs/PlaysAPI.md#publishplay) | **Post** /api/v1/plays/{name}/publish | Publish a play version
*PlaysAPI* | [**SavePlay**](docs/PlaysAPI.md#saveplay) | **Post** /api/v1/plays | Save a play (a new version)
*PlaysAPI* | [**SavePlaybook**](docs/PlaysAPI.md#saveplaybook) | **Post** /api/v1/playbooks | Save a playbook (pack)
*PlaysAPI* | [**SetPlaybookOn**](docs/PlaysAPI.md#setplaybookon) | **Post** /api/v1/playbooks/{name}/enable | Switch a pack on or off
*PlaysAPI* | [**StartDraft**](docs/PlaysAPI.md#startdraft) | **Post** /api/v1/drafts | Start a draft from a template
*PlaysAPI* | [**TestPlay**](docs/PlaysAPI.md#testplay) | **Post** /api/v1/plays/test | Check a play without saving it
*TeamsAPI* | [**ListTeams**](docs/TeamsAPI.md#listteams) | **Get** /api/v1/teams | Teams and their members
*TemplatesAPI* | [**CreateTemplateFromFile**](docs/TemplatesAPI.md#createtemplatefromfile) | **Post** /api/v1/templates/file | Make a template from a file with text tags
*TemplatesAPI* | [**CreateTemplateFromHtml**](docs/TemplatesAPI.md#createtemplatefromhtml) | **Post** /api/v1/templates/html | Make a template from HTML with text tags
*TemplatesAPI* | [**CreateTemplateFromSource**](docs/TemplatesAPI.md#createtemplatefromsource) | **Post** /api/v1/templates/source | Make a template from a document source
*TemplatesAPI* | [**CustomizeLibraryTemplate**](docs/TemplatesAPI.md#customizelibrarytemplate) | **Post** /api/v1/templates/library/{slug}/customize | Customize a built-in template
*TemplatesAPI* | [**DraftFromLibraryTemplate**](docs/TemplatesAPI.md#draftfromlibrarytemplate) | **Post** /api/v1/templates/library/{slug}/documents | Make a draft from a built-in template
*TemplatesAPI* | [**GetLibraryTemplate**](docs/TemplatesAPI.md#getlibrarytemplate) | **Get** /api/v1/templates/library/{slug} | One built-in template
*TemplatesAPI* | [**GetTemplate**](docs/TemplatesAPI.md#gettemplate) | **Get** /api/v1/templates/{id} | One template and its seats
*TemplatesAPI* | [**ListLibraryTemplates**](docs/TemplatesAPI.md#listlibrarytemplates) | **Get** /api/v1/templates/library | List the built-in templates
*TemplatesAPI* | [**SetTemplateDefaults**](docs/TemplatesAPI.md#settemplatedefaults) | **Post** /api/v1/templates/{id}/defaults | Set what a document made from a template starts with
*TemplatesAPI* | [**UseTemplate**](docs/TemplatesAPI.md#usetemplate) | **Post** /api/v1/templates/{id}/documents | Use a written template for people
*WebhooksAPI* | [**CreateWebhook**](docs/WebhooksAPI.md#createwebhook) | **Post** /api/v1/webhooks | Add a webhook
*WebhooksAPI* | [**DeleteWebhook**](docs/WebhooksAPI.md#deletewebhook) | **Delete** /api/v1/webhooks/{id} | Remove a webhook
*WebhooksAPI* | [**FetchRelayEvents**](docs/WebhooksAPI.md#fetchrelayevents) | **Get** /api/v1/webhooks/{id}/relay | Fetch stored deliveries of a relay webhook
*WebhooksAPI* | [**GetWebhook**](docs/WebhooksAPI.md#getwebhook) | **Get** /api/v1/webhooks/{id} | Get a webhook
*WebhooksAPI* | [**ListWebhookDeliveries**](docs/WebhooksAPI.md#listwebhookdeliveries) | **Get** /api/v1/webhooks/{id}/deliveries | Recent deliveries
*WebhooksAPI* | [**ListWebhookTopics**](docs/WebhooksAPI.md#listwebhooktopics) | **Get** /api/v1/webhooks/topics | The events a webhook can subscribe to
*WebhooksAPI* | [**ListWebhooks**](docs/WebhooksAPI.md#listwebhooks) | **Get** /api/v1/webhooks | List webhooks
*WebhooksAPI* | [**RotateWebhookSecret**](docs/WebhooksAPI.md#rotatewebhooksecret) | **Post** /api/v1/webhooks/{id}/rotate | Rotate the signing secret
*WebhooksAPI* | [**TestWebhook**](docs/WebhooksAPI.md#testwebhook) | **Post** /api/v1/webhooks/{id}/test | Send a test event


## Models

 - [AiDraftJob](docs/AiDraftJob.md)
 - [AiUsage](docs/AiUsage.md)
 - [AiUsageByFeatureValue](docs/AiUsageByFeatureValue.md)
 - [AnswerDraftRequest](docs/AnswerDraftRequest.md)
 - [BuilderSession](docs/BuilderSession.md)
 - [BuilderSessionRequest](docs/BuilderSessionRequest.md)
 - [CheckDocument200Response](docs/CheckDocument200Response.md)
 - [CheckDocumentRequest](docs/CheckDocumentRequest.md)
 - [CreateDocumentRequest](docs/CreateDocumentRequest.md)
 - [CreateTemplateFromSource201Response](docs/CreateTemplateFromSource201Response.md)
 - [CreateTemplateFromSourceRequest](docs/CreateTemplateFromSourceRequest.md)
 - [CreateWebhook201Response](docs/CreateWebhook201Response.md)
 - [CreateWebhookRequest](docs/CreateWebhookRequest.md)
 - [CustomizeLibraryTemplateRequest](docs/CustomizeLibraryTemplateRequest.md)
 - [DeleteDraft200Response](docs/DeleteDraft200Response.md)
 - [DeleteWebhook200Response](docs/DeleteWebhook200Response.md)
 - [DetectFieldsRequest](docs/DetectFieldsRequest.md)
 - [DetectFieldsResult](docs/DetectFieldsResult.md)
 - [DetectFieldsResultDocumentsInner](docs/DetectFieldsResultDocumentsInner.md)
 - [DetectFieldsResultPagesInner](docs/DetectFieldsResultPagesInner.md)
 - [DetectFieldsResultSummary](docs/DetectFieldsResultSummary.md)
 - [DetectedField](docs/DetectedField.md)
 - [Document](docs/Document.md)
 - [DocumentEventsInner](docs/DocumentEventsInner.md)
 - [DocumentMade](docs/DocumentMade.md)
 - [DocumentSummary](docs/DocumentSummary.md)
 - [DraftFromLibraryTemplateRequest](docs/DraftFromLibraryTemplateRequest.md)
 - [DraftWithAi202Response](docs/DraftWithAi202Response.md)
 - [DraftWithAiRequest](docs/DraftWithAiRequest.md)
 - [DuplicatePlayRequest](docs/DuplicatePlayRequest.md)
 - [EditWithAiRequest](docs/EditWithAiRequest.md)
 - [EditWithAiRequestVariablesInner](docs/EditWithAiRequestVariablesInner.md)
 - [EmailChoice](docs/EmailChoice.md)
 - [EmbedSession200Response](docs/EmbedSession200Response.md)
 - [EmbedSessionRequest](docs/EmbedSessionRequest.md)
 - [Error](docs/Error.md)
 - [ErrorError](docs/ErrorError.md)
 - [FetchRelayEvents200Response](docs/FetchRelayEvents200Response.md)
 - [FetchRelayEvents200ResponseEventsInner](docs/FetchRelayEvents200ResponseEventsInner.md)
 - [FileTemplateRequest](docs/FileTemplateRequest.md)
 - [GetDraftSource200Response](docs/GetDraftSource200Response.md)
 - [HtmlTemplateRequest](docs/HtmlTemplateRequest.md)
 - [ListDocuments200Response](docs/ListDocuments200Response.md)
 - [ListDocuments200ResponsePage](docs/ListDocuments200ResponsePage.md)
 - [ListEmails200Response](docs/ListEmails200Response.md)
 - [ListLibraryTemplates400Response](docs/ListLibraryTemplates400Response.md)
 - [ListLibraryTemplates400ResponseError](docs/ListLibraryTemplates400ResponseError.md)
 - [ListTeams200Response](docs/ListTeams200Response.md)
 - [ListTeams200ResponseTeamsInner](docs/ListTeams200ResponseTeamsInner.md)
 - [ListTeams200ResponseTeamsInnerMembersInner](docs/ListTeams200ResponseTeamsInnerMembersInner.md)
 - [ListTemplates200Response](docs/ListTemplates200Response.md)
 - [ListWebhookDeliveries200Response](docs/ListWebhookDeliveries200Response.md)
 - [ListWebhookDeliveries200ResponseDeliveriesInner](docs/ListWebhookDeliveries200ResponseDeliveriesInner.md)
 - [ListWebhookTopics200Response](docs/ListWebhookTopics200Response.md)
 - [ListWebhooks200Response](docs/ListWebhooks200Response.md)
 - [Party](docs/Party.md)
 - [Payment](docs/Payment.md)
 - [PaymentFee](docs/PaymentFee.md)
 - [PaymentsStatus200Response](docs/PaymentsStatus200Response.md)
 - [Person](docs/Person.md)
 - [PreflightAnswer](docs/PreflightAnswer.md)
 - [PreflightAnswerCost](docs/PreflightAnswerCost.md)
 - [PreflightAnswerPeriod](docs/PreflightAnswerPeriod.md)
 - [PreflightAnswerPlan](docs/PreflightAnswerPlan.md)
 - [PreflightOption](docs/PreflightOption.md)
 - [PreflightRequest](docs/PreflightRequest.md)
 - [PublishPlayRequest](docs/PublishPlayRequest.md)
 - [RefundPaymentRequest](docs/RefundPaymentRequest.md)
 - [Remind200Response](docs/Remind200Response.md)
 - [RemindRequest](docs/RemindRequest.md)
 - [ReviseDocument201Response](docs/ReviseDocument201Response.md)
 - [ReviseDocumentRequest](docs/ReviseDocumentRequest.md)
 - [RotateWebhookSecret200Response](docs/RotateWebhookSecret200Response.md)
 - [RotateWebhookSecretRequest](docs/RotateWebhookSecretRequest.md)
 - [SavePlayRequest](docs/SavePlayRequest.md)
 - [SavePlaybookRequest](docs/SavePlaybookRequest.md)
 - [ScribeSessionRequest](docs/ScribeSessionRequest.md)
 - [SendChoices](docs/SendChoices.md)
 - [SendDocument201Response](docs/SendDocument201Response.md)
 - [SendDocument201ResponseNotEmailedInner](docs/SendDocument201ResponseNotEmailedInner.md)
 - [SendDraft201Response](docs/SendDraft201Response.md)
 - [SendDraft202Response](docs/SendDraft202Response.md)
 - [SendDraft202ResponseApproval](docs/SendDraft202ResponseApproval.md)
 - [SendDraftRequest](docs/SendDraftRequest.md)
 - [SendRequest](docs/SendRequest.md)
 - [SendRequestPartiesInner](docs/SendRequestPartiesInner.md)
 - [SetPlaybookOnRequest](docs/SetPlaybookOnRequest.md)
 - [SetTemplateDefaults200Response](docs/SetTemplateDefaults200Response.md)
 - [SetTemplateDefaultsRequest](docs/SetTemplateDefaultsRequest.md)
 - [StartDraftRequest](docs/StartDraftRequest.md)
 - [StartDraftRequestAnswersValue](docs/StartDraftRequestAnswersValue.md)
 - [StartDraftRequestPeopleInner](docs/StartDraftRequestPeopleInner.md)
 - [Template](docs/Template.md)
 - [TemplateMade](docs/TemplateMade.md)
 - [TemplateMadeDetected](docs/TemplateMadeDetected.md)
 - [TemplateMadeWarningsInner](docs/TemplateMadeWarningsInner.md)
 - [TemplateSeatsInner](docs/TemplateSeatsInner.md)
 - [TestSign200Response](docs/TestSign200Response.md)
 - [TestSign200ResponseProgress](docs/TestSign200ResponseProgress.md)
 - [TestSignRequest](docs/TestSignRequest.md)
 - [TestWebhook200Response](docs/TestWebhook200Response.md)
 - [TestWebhookRequest](docs/TestWebhookRequest.md)
 - [Usage](docs/Usage.md)
 - [UsageBilling](docs/UsageBilling.md)
 - [UsageDocuments](docs/UsageDocuments.md)
 - [UsageDrafts](docs/UsageDrafts.md)
 - [UsageEmails](docs/UsageEmails.md)
 - [UsageLimits](docs/UsageLimits.md)
 - [UsagePeriod](docs/UsagePeriod.md)
 - [UsageSigners](docs/UsageSigners.md)
 - [UsageTemplates](docs/UsageTemplates.md)
 - [UseTemplateRequest](docs/UseTemplateRequest.md)
 - [VoidDocument200Response](docs/VoidDocument200Response.md)
 - [VoidDocumentRequest](docs/VoidDocumentRequest.md)
 - [WebhookDocumentsDocumentCreatedRequest](docs/WebhookDocumentsDocumentCreatedRequest.md)
 - [WebhookDocumentsDocumentSentRequest](docs/WebhookDocumentsDocumentSentRequest.md)
 - [WebhookEndpoint](docs/WebhookEndpoint.md)
 - [WizardBrief](docs/WizardBrief.md)
 - [WizardBriefPayment](docs/WizardBriefPayment.md)


## Authorization

Send your API key as a bearer token: `Authorization: Bearer $DOCUSTAY_API_KEY`. A test key (it starts with `ds_test_`) only ever emails you.

## Licence

MIT. See the LICENSE file. The Docustay server is AGPL-3.0; this client is MIT so you can ship it in any project.
