// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/runtime/security"
	"github.com/go-openapi/spec"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"go-server-server/awesomeProject/restapi/operations/item"
	"go-server-server/awesomeProject/restapi/operations/read_items"
)

// NewLostAndFoundAPI creates a new LostAndFound instance
func NewLostAndFoundAPI(spec *loads.Document) *LostAndFoundAPI {
	return &LostAndFoundAPI{
		handlers:            make(map[string]map[string]http.Handler),
		formats:             strfmt.Default,
		defaultConsumes:     "application/json",
		defaultProduces:     "application/json",
		customConsumers:     make(map[string]runtime.Consumer),
		customProducers:     make(map[string]runtime.Producer),
		PreServerShutdown:   func() {},
		ServerShutdown:      func() {},
		spec:                spec,
		useSwaggerUI:        false,
		ServeError:          errors.ServeError,
		BasicAuthenticator:  security.BasicAuth,
		APIKeyAuthenticator: security.APIKeyAuth,
		BearerAuthenticator: security.BearerAuth,

		JSONConsumer: runtime.JSONConsumer(),

		JSONProducer: runtime.JSONProducer(),

		ItemCreateItemHandler: item.CreateItemHandlerFunc(func(params item.CreateItemParams) middleware.Responder {
			return middleware.NotImplemented("operation item.CreateItem has not yet been implemented")
		}),
		ReadItemsGetItemByIDHandler: read_items.GetItemByIDHandlerFunc(func(params read_items.GetItemByIDParams) middleware.Responder {
			return middleware.NotImplemented("operation read_items.GetItemByID has not yet been implemented")
		}),
		ReadItemsListItemHandler: read_items.ListItemHandlerFunc(func(params read_items.ListItemParams) middleware.Responder {
			return middleware.NotImplemented("operation read_items.ListItem has not yet been implemented")
		}),
	}
}

/*LostAndFoundAPI the lost and found API */
type LostAndFoundAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	customConsumers map[string]runtime.Consumer
	customProducers map[string]runtime.Producer
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler
	useSwaggerUI    bool

	// BasicAuthenticator generates a runtime.Authenticator from the supplied basic auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator

	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator

	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator

	// JSONConsumer registers a consumer for the following mime types:
	//   - application/json
	JSONConsumer runtime.Consumer

	// JSONProducer registers a producer for the following mime types:
	//   - application/json
	JSONProducer runtime.Producer

	// ItemCreateItemHandler sets the operation handler for the create item operation
	ItemCreateItemHandler item.CreateItemHandler
	// ReadItemsGetItemByIDHandler sets the operation handler for the get item by Id operation
	ReadItemsGetItemByIDHandler read_items.GetItemByIDHandler
	// ReadItemsListItemHandler sets the operation handler for the list item operation
	ReadItemsListItemHandler read_items.ListItemHandler

	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// PreServerShutdown is called before the HTTP(S) server is shutdown
	// This allows for custom functions to get executed before the HTTP(S) server stops accepting traffic
	PreServerShutdown func()

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// UseRedoc for documentation at /docs
func (o *LostAndFoundAPI) UseRedoc() {
	o.useSwaggerUI = false
}

// UseSwaggerUI for documentation at /docs
func (o *LostAndFoundAPI) UseSwaggerUI() {
	o.useSwaggerUI = true
}

// SetDefaultProduces sets the default produces media type
func (o *LostAndFoundAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *LostAndFoundAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *LostAndFoundAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *LostAndFoundAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *LostAndFoundAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *LostAndFoundAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *LostAndFoundAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the LostAndFoundAPI
func (o *LostAndFoundAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.ItemCreateItemHandler == nil {
		unregistered = append(unregistered, "item.CreateItemHandler")
	}
	if o.ReadItemsGetItemByIDHandler == nil {
		unregistered = append(unregistered, "read_items.GetItemByIDHandler")
	}
	if o.ReadItemsListItemHandler == nil {
		unregistered = append(unregistered, "read_items.ListItemHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *LostAndFoundAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *LostAndFoundAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {
	return nil
}

// Authorizer returns the registered authorizer
func (o *LostAndFoundAPI) Authorizer() runtime.Authorizer {
	return nil
}

// ConsumersFor gets the consumers for the specified media types.
// MIME type parameters are ignored here.
func (o *LostAndFoundAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {
	result := make(map[string]runtime.Consumer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/json":
			result["application/json"] = o.JSONConsumer
		}

		if c, ok := o.customConsumers[mt]; ok {
			result[mt] = c
		}
	}
	return result
}

// ProducersFor gets the producers for the specified media types.
// MIME type parameters are ignored here.
func (o *LostAndFoundAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {
	result := make(map[string]runtime.Producer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/json":
			result["application/json"] = o.JSONProducer
		}

		if p, ok := o.customProducers[mt]; ok {
			result[mt] = p
		}
	}
	return result
}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *LostAndFoundAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	if path == "/" {
		path = ""
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the lost and found API
func (o *LostAndFoundAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *LostAndFoundAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened
	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/items"] = item.NewCreateItem(o.context, o.ItemCreateItemHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/items/{itemId}"] = read_items.NewGetItemByID(o.context, o.ReadItemsGetItemByIDHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/items"] = read_items.NewListItem(o.context, o.ReadItemsListItemHandler)
}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *LostAndFoundAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	if o.useSwaggerUI {
		return o.context.APIHandlerSwaggerUI(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middleware as you see fit
func (o *LostAndFoundAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

// RegisterConsumer allows you to add (or override) a consumer for a media type.
func (o *LostAndFoundAPI) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	o.customConsumers[mediaType] = consumer
}

// RegisterProducer allows you to add (or override) a producer for a media type.
func (o *LostAndFoundAPI) RegisterProducer(mediaType string, producer runtime.Producer) {
	o.customProducers[mediaType] = producer
}

// AddMiddlewareFor adds a http middleware to existing handler
func (o *LostAndFoundAPI) AddMiddlewareFor(method, path string, builder middleware.Builder) {
	um := strings.ToUpper(method)
	if path == "/" {
		path = ""
	}
	o.Init()
	if h, ok := o.handlers[um][path]; ok {
		o.handlers[method][path] = builder(h)
	}
}
