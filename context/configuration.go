package context

import "github.com/kataras/iris/v12/core/netutil"

// ConfigurationReadOnly can be implemented
// by Configuration, it's being used inside the Context.
// All methods that it contains should be "safe" to be called by the context
// at "serve time". A configuration field may be missing when it's not
// safe or its useless to be called from a request handler.
type ConfigurationReadOnly interface {
	// GetVHost returns the non-exported vhost config field.
	//
	// If original addr ended with :443 or :80, it will return the host without the port.
	// If original addr was :https or :http, it will return localhost.
	// If original addr was 0.0.0.0, it will return localhost.
	GetVHost() string

	// GetLogLevel returns the `Configuration.LogLevel` field.
	// The same (as `golog.LogLevel`) can be retrieved through `app.Logger().Level`.
	GetLogLevel() string

	// GetDisablePathCorrection returns the configuration.DisablePathCorrection,
	// DisablePathCorrection corrects and redirects the requested path to the registered path
	// for example, if /home/ path is requested but no handler for this Route found,
	// then the Router checks if /home handler exists, if yes,
	// (permant)redirects the client to the correct path /home.
	GetDisablePathCorrection() bool

	// GetDisablePathCorrectionRedirection returns the Configuration#DisablePathCorrectionRedirection field.
	// If DisablePathCorrectionRedirection set to true then it will handle paths as they are.
	// it will fire the handler of the matching route without
	// the last slash ("/") instead of send a redirection status.
	GetDisablePathCorrectionRedirection() bool

	// GetEnablePathEscape is the configuration.EnablePathEscape,
	// returns true when its escapes the path, the named parameters (if any).
	GetEnablePathEscape() bool

	// GetForceLowercaseRouting returns the value of the `ForceLowercaseRouting` setting.
	GetForceLowercaseRouting() bool

	// GetEnableOptimizations returns whether
	// the application has performance optimizations enabled.
	GetEnableOptimizations() bool

	// GetFireMethodNotAllowed returns the configuration.FireMethodNotAllowed.
	GetFireMethodNotAllowed() bool
	// GetDisableBodyConsumptionOnUnmarshal returns the configuration.GetDisableBodyConsumptionOnUnmarshal,
	// manages the reading behavior of the context's body readers/binders.
	// If returns true then the body consumption by the `context.UnmarshalBody/ReadJSON/ReadXML`
	// is disabled.
	//
	// By-default io.ReadAll` is used to read the body from the `context.Request.Body which is an `io.ReadCloser`,
	// if this field set to true then a new buffer will be created to read from and the request body.
	// The body will not be changed and existing data before the
	// context.UnmarshalBody/ReadJSON/ReadXML will be not consumed.
	GetDisableBodyConsumptionOnUnmarshal() bool

	// GetDisableAutoFireStatusCode returns the configuration.DisableAutoFireStatusCode.
	// Returns true when the http error status code handler automatic execution turned off.
	GetDisableAutoFireStatusCode() bool

	// GetTimeFormat returns the configuration.TimeFormat,
	// format for any kind of datetime parsing.
	GetTimeFormat() string

	// GetCharset returns the configuration.Charset,
	// the character encoding for various rendering
	// used for templates and the rest of the responses.
	GetCharset() string

	// GetPostMaxMemory returns the maximum configured post data size
	// that a client can send to the server, this differs
	// from the overral request body size which can be modified
	// by the `context#SetMaxRequestBodySize` or `iris#LimitRequestBodySize`.
	//
	// Defaults to 32MB or 32 << 20 if you prefer.
	GetPostMaxMemory() int64

	// GetTranslateLanguageContextKey returns the configuration's LocaleContextKey value,
	// used for i18n. Defaults to "iris.locale".
	GetLocaleContextKey() string
	// GetLanguageContextKey returns the configuration's LanguageContextKey value,
	// used for i18n. Defaults to "iris.locale.language".
	GetLanguageContextKey() string
	// GetVersionContextKey returns the configuration's VersionKey value,
	// used for API Versioning. Defaults to "iris.api.version".
	GetVersionContextKey() string
	// GetViewLayoutContextKey returns the key of the context's user values' key
	// which is being used to set the template
	// layout from a middleware or the main handler.
	// Overrides the parent's or the configuration's.
	GetViewLayoutContextKey() string
	// GetViewDataContextKey returns the key of the context's user values' key
	// which is being used to set the template
	// binding data from a middleware or the main handler.
	GetViewDataContextKey() string

	// GetRemoteAddrHeaders returns the allowed request headers names
	// that can be valid to parse the client's IP based on.
	//
	// Defaults to:
	// "X-Real-Ip":             true,
	// "X-Forwarded-For":       true,
	// "CF-Connecting-IP": false
	//
	// Look `context.RemoteAddr()` for more.
	GetRemoteAddrHeaders() map[string]bool
	// GetRemoteAddrPrivateSubnets returns the configuration's private sub-networks.
	// They are used to be compared against
	// IP Addresses fetched through `RemoteAddrHeaders` or `Request.RemoteAddr`.
	// For details please navigate through: https://github.com/kataras/iris/issues/1453
	// Defaults to an empty slice, usage:
	//
	// RemoteAddrPrivateSubnets {
	//	{Start: "10.0.0.0", End: "10.255.255.255"},
	//  {Start: "100.64.0.0", End: "100.127.255.255"},
	//	}
	//
	// Look `context.RemoteAddr()` for more.
	GetRemoteAddrPrivateSubnets() []netutil.IPRange
	// GetOther returns the configuration.Other map.
	GetOther() map[string]interface{}
}
