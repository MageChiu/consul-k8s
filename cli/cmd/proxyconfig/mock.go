package proxyconfig

const MockProxyConfig = `{
 "configs": [
  {
   "@type": "type.googleapis.com/envoy.admin.v3.BootstrapConfigDump",
   "bootstrap": {
    "node": {
     "id": "client-767ccfc8f9-bm82b-client-sidecar-proxy",
     "cluster": "client",
     "metadata": {
      "namespace": "default",
      "partition": "default"
     },
     "user_agent_name": "envoy",
     "user_agent_build_version": {
      "version": {
       "major_number": 1,
       "minor_number": 20,
       "patch": 2
      },
      "metadata": {
       "build.label": "dev",
       "ssl.version": "BoringSSL",
       "revision.status": "Clean",
       "revision.sha": "4aaf9593152c6996b9da384c8918e9ad4f0abd4d",
       "build.type": "RELEASE"
      }
     },
     "extensions": [
      {
       "name": "envoy.filters.dubbo.router",
       "category": "envoy.dubbo_proxy.filters"
      },
      {
       "name": "dubbo.hessian2",
       "category": "envoy.dubbo_proxy.serializers"
      },
      {
       "name": "dubbo",
       "category": "envoy.dubbo_proxy.protocols"
      },
      {
       "name": "composite-action",
       "category": "envoy.matching.action"
      },
      {
       "name": "skip",
       "category": "envoy.matching.action"
      },
      {
       "name": "default",
       "category": "envoy.dubbo_proxy.route_matchers"
      },
      {
       "name": "envoy.client_ssl_auth",
       "category": "envoy.filters.network"
      },
      {
       "name": "envoy.echo",
       "category": "envoy.filters.network"
      },
      {
       "name": "envoy.ext_authz",
       "category": "envoy.filters.network"
      },
      {
       "name": "envoy.filters.network.client_ssl_auth",
       "category": "envoy.filters.network"
      },
      {
       "name": "envoy.filters.network.connection_limit",
       "category": "envoy.filters.network"
      },
      {
       "name": "envoy.filters.network.direct_response",
       "category": "envoy.filters.network"
      },
      {
       "name": "envoy.filters.network.dubbo_proxy",
       "category": "envoy.filters.network"
      },
      {
       "name": "envoy.filters.network.echo",
       "category": "envoy.filters.network"
      },
      {
       "name": "envoy.filters.network.ext_authz",
       "category": "envoy.filters.network"
      },
      {
       "name": "envoy.filters.network.http_connection_manager",
       "category": "envoy.filters.network"
      },
      {
       "name": "envoy.filters.network.local_ratelimit",
       "category": "envoy.filters.network"
      },
      {
       "name": "envoy.filters.network.mongo_proxy",
       "category": "envoy.filters.network"
      },
      {
       "name": "envoy.filters.network.ratelimit",
       "category": "envoy.filters.network"
      },
      {
       "name": "envoy.filters.network.rbac",
       "category": "envoy.filters.network"
      },
      {
       "name": "envoy.filters.network.redis_proxy",
       "category": "envoy.filters.network"
      },
      {
       "name": "envoy.filters.network.sni_cluster",
       "category": "envoy.filters.network"
      },
      {
       "name": "envoy.filters.network.sni_dynamic_forward_proxy",
       "category": "envoy.filters.network"
      },
      {
       "name": "envoy.filters.network.tcp_proxy",
       "category": "envoy.filters.network"
      },
      {
       "name": "envoy.filters.network.thrift_proxy",
       "category": "envoy.filters.network"
      },
      {
       "name": "envoy.filters.network.wasm",
       "category": "envoy.filters.network"
      },
      {
       "name": "envoy.filters.network.zookeeper_proxy",
       "category": "envoy.filters.network"
      },
      {
       "name": "envoy.http_connection_manager",
       "category": "envoy.filters.network"
      },
      {
       "name": "envoy.mongo_proxy",
       "category": "envoy.filters.network"
      },
      {
       "name": "envoy.ratelimit",
       "category": "envoy.filters.network"
      },
      {
       "name": "envoy.redis_proxy",
       "category": "envoy.filters.network"
      },
      {
       "name": "envoy.tcp_proxy",
       "category": "envoy.filters.network"
      },
      {
       "name": "envoy.filters.listener.http_inspector",
       "category": "envoy.filters.listener"
      },
      {
       "name": "envoy.filters.listener.original_dst",
       "category": "envoy.filters.listener"
      },
      {
       "name": "envoy.filters.listener.original_src",
       "category": "envoy.filters.listener"
      },
      {
       "name": "envoy.filters.listener.proxy_protocol",
       "category": "envoy.filters.listener"
      },
      {
       "name": "envoy.filters.listener.tls_inspector",
       "category": "envoy.filters.listener"
      },
      {
       "name": "envoy.listener.http_inspector",
       "category": "envoy.filters.listener"
      },
      {
       "name": "envoy.listener.original_dst",
       "category": "envoy.filters.listener"
      },
      {
       "name": "envoy.listener.original_src",
       "category": "envoy.filters.listener"
      },
      {
       "name": "envoy.listener.proxy_protocol",
       "category": "envoy.filters.listener"
      },
      {
       "name": "envoy.listener.tls_inspector",
       "category": "envoy.filters.listener"
      },
      {
       "name": "envoy.transport_sockets.alts",
       "category": "envoy.transport_sockets.upstream"
      },
      {
       "name": "envoy.transport_sockets.quic",
       "category": "envoy.transport_sockets.upstream"
      },
      {
       "name": "envoy.transport_sockets.raw_buffer",
       "category": "envoy.transport_sockets.upstream"
      },
      {
       "name": "envoy.transport_sockets.starttls",
       "category": "envoy.transport_sockets.upstream"
      },
      {
       "name": "envoy.transport_sockets.tap",
       "category": "envoy.transport_sockets.upstream"
      },
      {
       "name": "envoy.transport_sockets.tls",
       "category": "envoy.transport_sockets.upstream"
      },
      {
       "name": "envoy.transport_sockets.upstream_proxy_protocol",
       "category": "envoy.transport_sockets.upstream"
      },
      {
       "name": "raw_buffer",
       "category": "envoy.transport_sockets.upstream"
      },
      {
       "name": "starttls",
       "category": "envoy.transport_sockets.upstream"
      },
      {
       "name": "tls",
       "category": "envoy.transport_sockets.upstream"
      },
      {
       "name": "envoy.filters.connection_pools.tcp.generic",
       "category": "envoy.upstreams"
      },
      {
       "name": "envoy.dynamic.ot",
       "category": "envoy.tracers"
      },
      {
       "name": "envoy.lightstep",
       "category": "envoy.tracers"
      },
      {
       "name": "envoy.tracers.datadog",
       "category": "envoy.tracers"
      },
      {
       "name": "envoy.tracers.dynamic_ot",
       "category": "envoy.tracers"
      },
      {
       "name": "envoy.tracers.lightstep",
       "category": "envoy.tracers"
      },
      {
       "name": "envoy.tracers.opencensus",
       "category": "envoy.tracers"
      },
      {
       "name": "envoy.tracers.skywalking",
       "category": "envoy.tracers"
      },
      {
       "name": "envoy.tracers.xray",
       "category": "envoy.tracers"
      },
      {
       "name": "envoy.tracers.zipkin",
       "category": "envoy.tracers"
      },
      {
       "name": "envoy.zipkin",
       "category": "envoy.tracers"
      },
      {
       "name": "envoy.resource_monitors.fixed_heap",
       "category": "envoy.resource_monitors"
      },
      {
       "name": "envoy.resource_monitors.injected_resource",
       "category": "envoy.resource_monitors"
      },
      {
       "name": "envoy.filters.thrift.rate_limit",
       "category": "envoy.thrift_proxy.filters"
      },
      {
       "name": "envoy.filters.thrift.router",
       "category": "envoy.thrift_proxy.filters"
      },
      {
       "name": "envoy.tls.cert_validator.default",
       "category": "envoy.tls.cert_validator"
      },
      {
       "name": "envoy.tls.cert_validator.spiffe",
       "category": "envoy.tls.cert_validator"
      },
      {
       "name": "envoy.internal_redirect_predicates.allow_listed_routes",
       "category": "envoy.internal_redirect_predicates"
      },
      {
       "name": "envoy.internal_redirect_predicates.previous_routes",
       "category": "envoy.internal_redirect_predicates"
      },
      {
       "name": "envoy.internal_redirect_predicates.safe_cross_scheme",
       "category": "envoy.internal_redirect_predicates"
      },
      {
       "name": "envoy.http.original_ip_detection.custom_header",
       "category": "envoy.http.original_ip_detection"
      },
      {
       "name": "envoy.http.original_ip_detection.xff",
       "category": "envoy.http.original_ip_detection"
      },
      {
       "name": "envoy.health_checkers.redis",
       "category": "envoy.health_checkers"
      },
      {
       "name": "envoy.request_id.uuid",
       "category": "envoy.request_id"
      },
      {
       "name": "envoy.retry_priorities.previous_priorities",
       "category": "envoy.retry_priorities"
      },
      {
       "name": "envoy.ip",
       "category": "envoy.resolvers"
      },
      {
       "name": "request-headers",
       "category": "envoy.matching.http.input"
      },
      {
       "name": "request-trailers",
       "category": "envoy.matching.http.input"
      },
      {
       "name": "response-headers",
       "category": "envoy.matching.http.input"
      },
      {
       "name": "response-trailers",
       "category": "envoy.matching.http.input"
      },
      {
       "name": "envoy.wasm.runtime.null",
       "category": "envoy.wasm.runtime"
      },
      {
       "name": "envoy.wasm.runtime.v8",
       "category": "envoy.wasm.runtime"
      },
      {
       "name": "envoy.quic.proof_source.filter_chain",
       "category": "envoy.quic.proof_source"
      },
      {
       "name": "envoy.transport_sockets.alts",
       "category": "envoy.transport_sockets.downstream"
      },
      {
       "name": "envoy.transport_sockets.quic",
       "category": "envoy.transport_sockets.downstream"
      },
      {
       "name": "envoy.transport_sockets.raw_buffer",
       "category": "envoy.transport_sockets.downstream"
      },
      {
       "name": "envoy.transport_sockets.starttls",
       "category": "envoy.transport_sockets.downstream"
      },
      {
       "name": "envoy.transport_sockets.tap",
       "category": "envoy.transport_sockets.downstream"
      },
      {
       "name": "envoy.transport_sockets.tls",
       "category": "envoy.transport_sockets.downstream"
      },
      {
       "name": "raw_buffer",
       "category": "envoy.transport_sockets.downstream"
      },
      {
       "name": "starttls",
       "category": "envoy.transport_sockets.downstream"
      },
      {
       "name": "tls",
       "category": "envoy.transport_sockets.downstream"
      },
      {
       "name": "envoy.bootstrap.wasm",
       "category": "envoy.bootstrap"
      },
      {
       "name": "envoy.extensions.network.socket_interface.default_socket_interface",
       "category": "envoy.bootstrap"
      },
      {
       "name": "envoy.filters.udp.dns_filter",
       "category": "envoy.filters.udp_listener"
      },
      {
       "name": "envoy.filters.udp_listener.udp_proxy",
       "category": "envoy.filters.udp_listener"
      },
      {
       "name": "envoy.extensions.upstreams.http.v3.HttpProtocolOptions",
       "category": "envoy.upstream_options"
      },
      {
       "name": "envoy.upstreams.http.http_protocol_options",
       "category": "envoy.upstream_options"
      },
      {
       "name": "envoy.compression.brotli.compressor",
       "category": "envoy.compression.compressor"
      },
      {
       "name": "envoy.compression.gzip.compressor",
       "category": "envoy.compression.compressor"
      },
      {
       "name": "envoy.cluster.eds",
       "category": "envoy.clusters"
      },
      {
       "name": "envoy.cluster.logical_dns",
       "category": "envoy.clusters"
      },
      {
       "name": "envoy.cluster.original_dst",
       "category": "envoy.clusters"
      },
      {
       "name": "envoy.cluster.static",
       "category": "envoy.clusters"
      },
      {
       "name": "envoy.cluster.strict_dns",
       "category": "envoy.clusters"
      },
      {
       "name": "envoy.clusters.aggregate",
       "category": "envoy.clusters"
      },
      {
       "name": "envoy.clusters.dynamic_forward_proxy",
       "category": "envoy.clusters"
      },
      {
       "name": "envoy.clusters.redis",
       "category": "envoy.clusters"
      },
      {
       "name": "envoy.access_loggers.file",
       "category": "envoy.access_loggers"
      },
      {
       "name": "envoy.access_loggers.http_grpc",
       "category": "envoy.access_loggers"
      },
      {
       "name": "envoy.access_loggers.open_telemetry",
       "category": "envoy.access_loggers"
      },
      {
       "name": "envoy.access_loggers.stderr",
       "category": "envoy.access_loggers"
      },
      {
       "name": "envoy.access_loggers.stdout",
       "category": "envoy.access_loggers"
      },
      {
       "name": "envoy.access_loggers.tcp_grpc",
       "category": "envoy.access_loggers"
      },
      {
       "name": "envoy.access_loggers.wasm",
       "category": "envoy.access_loggers"
      },
      {
       "name": "envoy.file_access_log",
       "category": "envoy.access_loggers"
      },
      {
       "name": "envoy.http_grpc_access_log",
       "category": "envoy.access_loggers"
      },
      {
       "name": "envoy.open_telemetry_access_log",
       "category": "envoy.access_loggers"
      },
      {
       "name": "envoy.stderr_access_log",
       "category": "envoy.access_loggers"
      },
      {
       "name": "envoy.stdout_access_log",
       "category": "envoy.access_loggers"
      },
      {
       "name": "envoy.tcp_grpc_access_log",
       "category": "envoy.access_loggers"
      },
      {
       "name": "envoy.wasm_access_log",
       "category": "envoy.access_loggers"
      },
      {
       "name": "envoy.rbac.matchers.upstream.upstream_ip_port",
       "category": "envoy.rbac.matchers"
      },
      {
       "name": "auto",
       "category": "envoy.thrift_proxy.protocols"
      },
      {
       "name": "binary",
       "category": "envoy.thrift_proxy.protocols"
      },
      {
       "name": "binary/non-strict",
       "category": "envoy.thrift_proxy.protocols"
      },
      {
       "name": "compact",
       "category": "envoy.thrift_proxy.protocols"
      },
      {
       "name": "twitter",
       "category": "envoy.thrift_proxy.protocols"
      },
      {
       "name": "envoy.rate_limit_descriptors.expr",
       "category": "envoy.rate_limit_descriptors"
      },
      {
       "name": "envoy.matching.matchers.consistent_hashing",
       "category": "envoy.matching.input_matchers"
      },
      {
       "name": "envoy.matching.matchers.ip",
       "category": "envoy.matching.input_matchers"
      },
      {
       "name": "envoy.watchdog.abort_action",
       "category": "envoy.guarddog_actions"
      },
      {
       "name": "envoy.watchdog.profile_action",
       "category": "envoy.guarddog_actions"
      },
      {
       "name": "envoy.quic.crypto_stream.server.quiche",
       "category": "envoy.quic.server.crypto_stream"
      },
      {
       "name": "envoy.key_value.file_based",
       "category": "envoy.common.key_value"
      },
      {
       "name": "envoy.bandwidth_limit",
       "category": "envoy.filters.http"
      },
      {
       "name": "envoy.buffer",
       "category": "envoy.filters.http"
      },
      {
       "name": "envoy.cors",
       "category": "envoy.filters.http"
      },
      {
       "name": "envoy.csrf",
       "category": "envoy.filters.http"
      },
      {
       "name": "envoy.ext_authz",
       "category": "envoy.filters.http"
      },
      {
       "name": "envoy.ext_proc",
       "category": "envoy.filters.http"
      },
      {
       "name": "envoy.fault",
       "category": "envoy.filters.http"
      },
      {
       "name": "envoy.filters.http.adaptive_concurrency",
       "category": "envoy.filters.http"
      },
      {
       "name": "envoy.filters.http.admission_control",
       "category": "envoy.filters.http"
      },
      {
       "name": "envoy.filters.http.alternate_protocols_cache",
       "category": "envoy.filters.http"
      },
      {
       "name": "envoy.filters.http.aws_lambda",
       "category": "envoy.filters.http"
      },
      {
       "name": "envoy.filters.http.aws_request_signing",
       "category": "envoy.filters.http"
      },
      {
       "name": "envoy.filters.http.bandwidth_limit",
       "category": "envoy.filters.http"
      },
      {
       "name": "envoy.filters.http.buffer",
       "category": "envoy.filters.http"
      },
      {
       "name": "envoy.filters.http.cache",
       "category": "envoy.filters.http"
      },
      {
       "name": "envoy.filters.http.cdn_loop",
       "category": "envoy.filters.http"
      },
      {
       "name": "envoy.filters.http.composite",
       "category": "envoy.filters.http"
      },
      {
       "name": "envoy.filters.http.compressor",
       "category": "envoy.filters.http"
      },
      {
       "name": "envoy.filters.http.cors",
       "category": "envoy.filters.http"
      },
      {
       "name": "envoy.filters.http.csrf",
       "category": "envoy.filters.http"
      },
      {
       "name": "envoy.filters.http.decompressor",
       "category": "envoy.filters.http"
      },
      {
       "name": "envoy.filters.http.dynamic_forward_proxy",
       "category": "envoy.filters.http"
      },
      {
       "name": "envoy.filters.http.dynamo",
       "category": "envoy.filters.http"
      },
      {
       "name": "envoy.filters.http.ext_authz",
       "category": "envoy.filters.http"
      },
      {
       "name": "envoy.filters.http.ext_proc",
       "category": "envoy.filters.http"
      },
      {
       "name": "envoy.filters.http.fault",
       "category": "envoy.filters.http"
      },
      {
       "name": "envoy.filters.http.grpc_http1_bridge",
       "category": "envoy.filters.http"
      },
      {
       "name": "envoy.filters.http.grpc_http1_reverse_bridge",
       "category": "envoy.filters.http"
      },
      {
       "name": "envoy.filters.http.grpc_json_transcoder",
       "category": "envoy.filters.http"
      },
      {
       "name": "envoy.filters.http.grpc_stats",
       "category": "envoy.filters.http"
      },
      {
       "name": "envoy.filters.http.grpc_web",
       "category": "envoy.filters.http"
      },
      {
       "name": "envoy.filters.http.header_to_metadata",
       "category": "envoy.filters.http"
      },
      {
       "name": "envoy.filters.http.health_check",
       "category": "envoy.filters.http"
      },
      {
       "name": "envoy.filters.http.ip_tagging",
       "category": "envoy.filters.http"
      },
      {
       "name": "envoy.filters.http.jwt_authn",
       "category": "envoy.filters.http"
      },
      {
       "name": "envoy.filters.http.local_ratelimit",
       "category": "envoy.filters.http"
      },
      {
       "name": "envoy.filters.http.lua",
       "category": "envoy.filters.http"
      },
      {
       "name": "envoy.filters.http.oauth2",
       "category": "envoy.filters.http"
      },
      {
       "name": "envoy.filters.http.on_demand",
       "category": "envoy.filters.http"
      },
      {
       "name": "envoy.filters.http.original_src",
       "category": "envoy.filters.http"
      },
      {
       "name": "envoy.filters.http.ratelimit",
       "category": "envoy.filters.http"
      },
      {
       "name": "envoy.filters.http.rbac",
       "category": "envoy.filters.http"
      },
      {
       "name": "envoy.filters.http.router",
       "category": "envoy.filters.http"
      },
      {
       "name": "envoy.filters.http.set_metadata",
       "category": "envoy.filters.http"
      },
      {
       "name": "envoy.filters.http.tap",
       "category": "envoy.filters.http"
      },
      {
       "name": "envoy.filters.http.wasm",
       "category": "envoy.filters.http"
      },
      {
       "name": "envoy.grpc_http1_bridge",
       "category": "envoy.filters.http"
      },
      {
       "name": "envoy.grpc_json_transcoder",
       "category": "envoy.filters.http"
      },
      {
       "name": "envoy.grpc_web",
       "category": "envoy.filters.http"
      },
      {
       "name": "envoy.health_check",
       "category": "envoy.filters.http"
      },
      {
       "name": "envoy.http_dynamo_filter",
       "category": "envoy.filters.http"
      },
      {
       "name": "envoy.ip_tagging",
       "category": "envoy.filters.http"
      },
      {
       "name": "envoy.local_rate_limit",
       "category": "envoy.filters.http"
      },
      {
       "name": "envoy.lua",
       "category": "envoy.filters.http"
      },
      {
       "name": "envoy.rate_limit",
       "category": "envoy.filters.http"
      },
      {
       "name": "envoy.router",
       "category": "envoy.filters.http"
      },
      {
       "name": "match-wrapper",
       "category": "envoy.filters.http"
      },
      {
       "name": "envoy.compression.brotli.decompressor",
       "category": "envoy.compression.decompressor"
      },
      {
       "name": "envoy.compression.gzip.decompressor",
       "category": "envoy.compression.decompressor"
      },
      {
       "name": "envoy.matching.common_inputs.environment_variable",
       "category": "envoy.matching.common_inputs"
      },
      {
       "name": "envoy.dog_statsd",
       "category": "envoy.stats_sinks"
      },
      {
       "name": "envoy.graphite_statsd",
       "category": "envoy.stats_sinks"
      },
      {
       "name": "envoy.metrics_service",
       "category": "envoy.stats_sinks"
      },
      {
       "name": "envoy.stat_sinks.dog_statsd",
       "category": "envoy.stats_sinks"
      },
      {
       "name": "envoy.stat_sinks.graphite_statsd",
       "category": "envoy.stats_sinks"
      },
      {
       "name": "envoy.stat_sinks.hystrix",
       "category": "envoy.stats_sinks"
      },
      {
       "name": "envoy.stat_sinks.metrics_service",
       "category": "envoy.stats_sinks"
      },
      {
       "name": "envoy.stat_sinks.statsd",
       "category": "envoy.stats_sinks"
      },
      {
       "name": "envoy.stat_sinks.wasm",
       "category": "envoy.stats_sinks"
      },
      {
       "name": "envoy.statsd",
       "category": "envoy.stats_sinks"
      },
      {
       "name": "envoy.retry_host_predicates.omit_canary_hosts",
       "category": "envoy.retry_host_predicates"
      },
      {
       "name": "envoy.retry_host_predicates.omit_host_metadata",
       "category": "envoy.retry_host_predicates"
      },
      {
       "name": "envoy.retry_host_predicates.previous_hosts",
       "category": "envoy.retry_host_predicates"
      },
      {
       "name": "auto",
       "category": "envoy.thrift_proxy.transports"
      },
      {
       "name": "framed",
       "category": "envoy.thrift_proxy.transports"
      },
      {
       "name": "header",
       "category": "envoy.thrift_proxy.transports"
      },
      {
       "name": "unframed",
       "category": "envoy.thrift_proxy.transports"
      },
      {
       "name": "preserve_case",
       "category": "envoy.http.stateful_header_formatters"
      },
      {
       "name": "envoy.grpc_credentials.aws_iam",
       "category": "envoy.grpc_credentials"
      },
      {
       "name": "envoy.grpc_credentials.default",
       "category": "envoy.grpc_credentials"
      },
      {
       "name": "envoy.grpc_credentials.file_based_metadata",
       "category": "envoy.grpc_credentials"
      },
      {
       "name": "envoy.extensions.http.cache.simple",
       "category": "envoy.http.cache"
      },
      {
       "name": "envoy.formatter.metadata",
       "category": "envoy.formatter"
      },
      {
       "name": "envoy.formatter.req_without_query",
       "category": "envoy.formatter"
      }
     ]
    },
    "static_resources": {
     "clusters": [
      {
       "name": "local_agent",
       "type": "STATIC",
       "connect_timeout": "1s",
       "http2_protocol_options": {},
       "load_assignment": {
        "cluster_name": "local_agent",
        "endpoints": [
         {
          "lb_endpoints": [
           {
            "endpoint": {
             "address": {
              "socket_address": {
               "address": "172.18.0.2",
               "port_value": 8502
              }
             }
            }
           }
          ]
         }
        ]
       }
      }
     ]
    },
    "dynamic_resources": {
     "lds_config": {
      "ads": {},
      "resource_api_version": "V3"
     },
     "cds_config": {
      "ads": {},
      "resource_api_version": "V3"
     },
     "ads_config": {
      "api_type": "DELTA_GRPC",
      "grpc_services": [
       {
        "envoy_grpc": {
         "cluster_name": "local_agent"
        },
        "initial_metadata": [
         {
          "key": "x-consul-token"
         }
        ]
       }
      ],
      "transport_api_version": "V3"
     }
    },
    "admin": {
     "access_log_path": "/dev/null",
     "address": {
      "socket_address": {
       "address": "127.0.0.1",
       "port_value": 19000
      }
     }
    },
    "stats_config": {
     "stats_tags": [
      {
       "tag_name": "consul.destination.custom_hash",
       "regex": "^cluster\\.(?:passthrough~)?((?:([^.]+)~)?(?:[^.]+\\.)?[^.]+\\.[^.]+\\.(?:[^.]+\\.)?[^.]+\\.[^.]+\\.[^.]+\\.consul\\.)"
      },
      {
       "tag_name": "consul.destination.service_subset",
       "regex": "^cluster\\.(?:passthrough~)?((?:[^.]+~)?(?:([^.]+)\\.)?[^.]+\\.[^.]+\\.(?:[^.]+\\.)?[^.]+\\.[^.]+\\.[^.]+\\.consul\\.)"
      },
      {
       "tag_name": "consul.destination.service",
       "regex": "^cluster\\.(?:passthrough~)?((?:[^.]+~)?(?:[^.]+\\.)?([^.]+)\\.[^.]+\\.(?:[^.]+\\.)?[^.]+\\.[^.]+\\.[^.]+\\.consul\\.)"
      },
      {
       "tag_name": "consul.destination.namespace",
       "regex": "^cluster\\.(?:passthrough~)?((?:[^.]+~)?(?:[^.]+\\.)?[^.]+\\.([^.]+)\\.(?:[^.]+\\.)?[^.]+\\.[^.]+\\.[^.]+\\.consul\\.)"
      },
      {
       "tag_name": "consul.destination.partition",
       "regex": "^cluster\\.(?:passthrough~)?((?:[^.]+~)?(?:[^.]+\\.)?[^.]+\\.[^.]+\\.(?:([^.]+)\\.)?[^.]+\\.[^.]+\\.[^.]+\\.consul\\.)"
      },
      {
       "tag_name": "consul.destination.datacenter",
       "regex": "^cluster\\.(?:passthrough~)?((?:[^.]+~)?(?:[^.]+\\.)?[^.]+\\.[^.]+\\.(?:[^.]+\\.)?([^.]+)\\.[^.]+\\.[^.]+\\.consul\\.)"
      },
      {
       "tag_name": "consul.destination.routing_type",
       "regex": "^cluster\\.(?:passthrough~)?((?:[^.]+~)?(?:[^.]+\\.)?[^.]+\\.[^.]+\\.(?:[^.]+\\.)?[^.]+\\.([^.]+)\\.[^.]+\\.consul\\.)"
      },
      {
       "tag_name": "consul.destination.trust_domain",
       "regex": "^cluster\\.(?:passthrough~)?((?:[^.]+~)?(?:[^.]+\\.)?[^.]+\\.[^.]+\\.(?:[^.]+\\.)?[^.]+\\.[^.]+\\.([^.]+)\\.consul\\.)"
      },
      {
       "tag_name": "consul.destination.target",
       "regex": "^cluster\\.(?:passthrough~)?(((?:[^.]+~)?(?:[^.]+\\.)?[^.]+\\.[^.]+\\.(?:[^.]+\\.)?[^.]+)\\.[^.]+\\.[^.]+\\.consul\\.)"
      },
      {
       "tag_name": "consul.destination.full_target",
       "regex": "^cluster\\.(?:passthrough~)?(((?:[^.]+~)?(?:[^.]+\\.)?[^.]+\\.[^.]+\\.(?:[^.]+\\.)?[^.]+\\.[^.]+\\.[^.]+)\\.consul\\.)"
      },
      {
       "tag_name": "consul.upstream.service",
       "regex": "^(?:tcp|http)\\.upstream\\.(([^.]+)(?:\\.[^.]+)?(?:\\.[^.]+)?\\.[^.]+\\.)"
      },
      {
       "tag_name": "consul.upstream.datacenter",
       "regex": "^(?:tcp|http)\\.upstream\\.([^.]+(?:\\.[^.]+)?(?:\\.[^.]+)?\\.([^.]+)\\.)"
      },
      {
       "tag_name": "consul.upstream.namespace",
       "regex": "^(?:tcp|http)\\.upstream\\.([^.]+(?:\\.([^.]+))?(?:\\.[^.]+)?\\.[^.]+\\.)"
      },
      {
       "tag_name": "consul.upstream.partition",
       "regex": "^(?:tcp|http)\\.upstream\\.([^.]+(?:\\.[^.]+)?(?:\\.([^.]+))?\\.[^.]+\\.)"
      },
      {
       "tag_name": "consul.custom_hash",
       "regex": "^cluster\\.((?:([^.]+)~)?(?:[^.]+\\.)?[^.]+\\.[^.]+\\.(?:[^.]+\\.)?[^.]+\\.[^.]+\\.[^.]+\\.consul\\.)"
      },
      {
       "tag_name": "consul.service_subset",
       "regex": "^cluster\\.((?:[^.]+~)?(?:([^.]+)\\.)?[^.]+\\.[^.]+\\.(?:[^.]+\\.)?[^.]+\\.[^.]+\\.[^.]+\\.consul\\.)"
      },
      {
       "tag_name": "consul.service",
       "regex": "^cluster\\.((?:[^.]+~)?(?:[^.]+\\.)?([^.]+)\\.[^.]+\\.(?:[^.]+\\.)?[^.]+\\.[^.]+\\.[^.]+\\.consul\\.)"
      },
      {
       "tag_name": "consul.namespace",
       "regex": "^cluster\\.((?:[^.]+~)?(?:[^.]+\\.)?[^.]+\\.([^.]+)\\.(?:[^.]+\\.)?[^.]+\\.[^.]+\\.[^.]+\\.consul\\.)"
      },
      {
       "tag_name": "consul.datacenter",
       "regex": "^cluster\\.((?:[^.]+~)?(?:[^.]+\\.)?[^.]+\\.[^.]+\\.(?:[^.]+\\.)?([^.]+)\\.[^.]+\\.[^.]+\\.consul\\.)"
      },
      {
       "tag_name": "consul.routing_type",
       "regex": "^cluster\\.((?:[^.]+~)?(?:[^.]+\\.)?[^.]+\\.[^.]+\\.(?:[^.]+\\.)?[^.]+\\.([^.]+)\\.[^.]+\\.consul\\.)"
      },
      {
       "tag_name": "consul.trust_domain",
       "regex": "^cluster\\.((?:[^.]+~)?(?:[^.]+\\.)?[^.]+\\.[^.]+\\.(?:[^.]+\\.)?[^.]+\\.[^.]+\\.([^.]+)\\.consul\\.)"
      },
      {
       "tag_name": "consul.target",
       "regex": "^cluster\\.(((?:[^.]+~)?(?:[^.]+\\.)?[^.]+\\.[^.]+\\.(?:[^.]+\\.)?[^.]+)\\.[^.]+\\.[^.]+\\.consul\\.)"
      },
      {
       "tag_name": "consul.full_target",
       "regex": "^cluster\\.(((?:[^.]+~)?(?:[^.]+\\.)?[^.]+\\.[^.]+\\.(?:[^.]+\\.)?[^.]+\\.[^.]+\\.[^.]+)\\.consul\\.)"
      },
      {
       "tag_name": "local_cluster",
       "fixed_value": "client"
      },
      {
       "tag_name": "consul.source.service",
       "fixed_value": "client"
      },
      {
       "tag_name": "consul.source.namespace",
       "fixed_value": "default"
      },
      {
       "tag_name": "consul.source.partition",
       "fixed_value": "default"
      },
      {
       "tag_name": "consul.source.datacenter",
       "fixed_value": "dc1"
      }
     ],
     "use_all_default_tags": true
    }
   },
   "last_updated": "2022-03-01T22:57:32.720Z"
  },
  {
   "@type": "type.googleapis.com/envoy.admin.v3.ClustersConfigDump",
   "static_clusters": [
    {
     "cluster": {
      "@type": "type.googleapis.com/envoy.config.cluster.v3.Cluster",
      "name": "local_agent",
      "type": "STATIC",
      "connect_timeout": "1s",
      "http2_protocol_options": {},
      "load_assignment": {
       "cluster_name": "local_agent",
       "endpoints": [
        {
         "lb_endpoints": [
          {
           "endpoint": {
            "address": {
             "socket_address": {
              "address": "172.18.0.2",
              "port_value": 8502
             }
            }
           }
          }
         ]
        }
       ]
      }
     },
     "last_updated": "2022-03-01T22:57:32.772Z"
    }
   ],
   "dynamic_active_clusters": [
    {
     "version_info": "ad70d0815091ccaae13efcd519d29aa591ba7d51043f6ee692d960555ee81cd2",
     "cluster": {
      "@type": "type.googleapis.com/envoy.config.cluster.v3.Cluster",
      "name": "local_app",
      "type": "STATIC",
      "connect_timeout": "5s",
      "load_assignment": {
       "cluster_name": "local_app",
       "endpoints": [
        {
         "lb_endpoints": [
          {
           "endpoint": {
            "address": {
             "socket_address": {
              "address": "127.0.0.1",
              "port_value": 0
             }
            }
           }
          }
         ]
        }
       ]
      }
     },
     "last_updated": "2022-03-01T22:57:32.929Z"
    },
    {
     "version_info": "dcd8d0247ba0149bfdc151428353b3f29d0665bf5c12af6a105a0abcc5af40ac",
     "cluster": {
      "@type": "type.googleapis.com/envoy.config.cluster.v3.Cluster",
      "name": "original-destination",
      "type": "ORIGINAL_DST",
      "connect_timeout": "5s",
      "lb_policy": "CLUSTER_PROVIDED"
     },
     "last_updated": "2022-03-01T22:57:32.956Z"
    },
    {
     "version_info": "bda39a9054edb9c767486b7837e11e5e0f28551db2d0c7d31c3be4955fc04f50",
     "cluster": {
      "@type": "type.googleapis.com/envoy.config.cluster.v3.Cluster",
      "name": "server.default.dc1.internal.4563c051-0a92-e021-da69-0be9acd651da.consul",
      "type": "EDS",
      "eds_cluster_config": {
       "eds_config": {
        "ads": {},
        "resource_api_version": "V3"
       }
      },
      "connect_timeout": "5s",
      "circuit_breakers": {},
      "outlier_detection": {},
      "transport_socket": {
       "name": "tls",
       "typed_config": {
        "@type": "type.googleapis.com/envoy.extensions.transport_sockets.tls.v3.UpstreamTlsContext",
        "common_tls_context": {
         "tls_params": {},
         "tls_certificates": [
          {
           "certificate_chain": {
            "inline_string": "-----BEGIN CERTIFICATE-----\nMIICGTCCAb+gAwIBAgIBCTAKBggqhkjOPQQDAjAwMS4wLAYDVQQDEyVwcmktYWlz\nY3dnMS5jb25zdWwuY2EuNDU2M2MwNTEuY29uc3VsMB4XDTIyMDMwMTIyNTYxOFoX\nDTIyMDMwNDIyNTYxOFowADBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABOkztqwq\nP4SnSZ+T1JIakPeSrgcL+k30wu7rAE+xVN5lsY+iK6DAIVmHapLkOsuElI13arJa\nDaaqqdWJUG2LtqGjgfkwgfYwDgYDVR0PAQH/BAQDAgO4MB0GA1UdJQQWMBQGCCsG\nAQUFBwMCBggrBgEFBQcDATAMBgNVHRMBAf8EAjAAMCkGA1UdDgQiBCB7wCCVHVTd\nv6C07SflIf2lX1pvC1wlQIQi2zrhxaBg7TArBgNVHSMEJDAigCAMfL0aTpEwCQMh\nrD6OZMrC7lJyKSB339GwDGyU4OV3vzBfBgNVHREBAf8EVTBThlFzcGlmZmU6Ly80\nNTYzYzA1MS0wYTkyLWUwMjEtZGE2OS0wYmU5YWNkNjUxZGEuY29uc3VsL25zL2Rl\nZmF1bHQvZGMvZGMxL3N2Yy9jbGllbnQwCgYIKoZIzj0EAwIDSAAwRQIhAKKrhL0B\ny4PR/8a30JC7BmBmNWxrPSRIBaLsdhMJ9CDPAiAA7RJqkh1sc6XLx65P9FYSqDxT\nViilKSWGfQ23Ik8i1Q==\n-----END CERTIFICATE-----\n"
           },
           "private_key": {
            "inline_string": "[redacted]"
           }
          }
         ],
         "validation_context": {
          "trusted_ca": {
           "inline_string": "-----BEGIN CERTIFICATE-----\nMIICDjCCAbOgAwIBAgIBBzAKBggqhkjOPQQDAjAwMS4wLAYDVQQDEyVwcmktYWlz\nY3dnMS5jb25zdWwuY2EuNDU2M2MwNTEuY29uc3VsMB4XDTIyMDMwMTIyNTY1OFoX\nDTMyMDIyNzIyNTY1OFowMDEuMCwGA1UEAxMlcHJpLWFpc2N3ZzEuY29uc3VsLmNh\nLjQ1NjNjMDUxLmNvbnN1bDBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABDzGo4Ao\nas3SpBZl+0/WG8MWcuMgcu/VihHmxs+kRlepVWC+H9KA0IwtUTKgVtCKf7qp7dbJ\nG54R4tgv5qG6X/Cjgb0wgbowDgYDVR0PAQH/BAQDAgGGMA8GA1UdEwEB/wQFMAMB\nAf8wKQYDVR0OBCIEIAx8vRpOkTAJAyGsPo5kysLuUnIpIHff0bAMbJTg5Xe/MCsG\nA1UdIwQkMCKAIAx8vRpOkTAJAyGsPo5kysLuUnIpIHff0bAMbJTg5Xe/MD8GA1Ud\nEQQ4MDaGNHNwaWZmZTovLzQ1NjNjMDUxLTBhOTItZTAyMS1kYTY5LTBiZTlhY2Q2\nNTFkYS5jb25zdWwwCgYIKoZIzj0EAwIDSQAwRgIhALkoAuOMTRSTMQnByTCN11Uk\nOsd9eet4efD8tJtXwXppAiEAl6Fd/cWtGLT3ciEOpgxJMIEBwTKtd9xO/KJa67Cr\nZJM=\n-----END CERTIFICATE-----\n"
          },
          "match_subject_alt_names": [
           {
            "exact": "spiffe://4563c051-0a92-e021-da69-0be9acd651da.consul/ns/default/dc/dc1/svc/server"
           }
          ]
         }
        },
        "sni": "server.default.dc1.internal.4563c051-0a92-e021-da69-0be9acd651da.consul"
       }
      },
      "common_lb_config": {
       "healthy_panic_threshold": {}
      },
      "alt_stat_name": "server.default.dc1.internal.4563c051-0a92-e021-da69-0be9acd651da.consul"
     },
     "last_updated": "2022-03-01T22:57:32.908Z"
    }
   ]
  },
  {
   "@type": "type.googleapis.com/envoy.admin.v3.ListenersConfigDump",
   "dynamic_listeners": [
    {
     "name": "public_listener:10.244.0.51:20000",
     "active_state": {
      "version_info": "2d89771a85b8046d3c0e0e045ed89a24611c2b047d141a032b199a1bfbd682d2",
      "listener": {
       "@type": "type.googleapis.com/envoy.config.listener.v3.Listener",
       "name": "public_listener:10.244.0.51:20000",
       "address": {
        "socket_address": {
         "address": "10.244.0.51",
         "port_value": 20000
        }
       },
       "filter_chains": [
        {
         "filters": [
          {
           "name": "envoy.filters.network.rbac",
           "typed_config": {
            "@type": "type.googleapis.com/envoy.extensions.filters.network.rbac.v3.RBAC",
            "rules": {
             "action": "DENY"
            },
            "stat_prefix": "connect_authz"
           }
          },
          {
           "name": "envoy.filters.network.tcp_proxy",
           "typed_config": {
            "@type": "type.googleapis.com/envoy.extensions.filters.network.tcp_proxy.v3.TcpProxy",
            "stat_prefix": "public_listener",
            "cluster": "local_app"
           }
          }
         ],
         "transport_socket": {
          "name": "tls",
          "typed_config": {
           "@type": "type.googleapis.com/envoy.extensions.transport_sockets.tls.v3.DownstreamTlsContext",
           "common_tls_context": {
            "tls_params": {},
            "tls_certificates": [
             {
              "certificate_chain": {
               "inline_string": "-----BEGIN CERTIFICATE-----\nMIICGTCCAb+gAwIBAgIBCTAKBggqhkjOPQQDAjAwMS4wLAYDVQQDEyVwcmktYWlz\nY3dnMS5jb25zdWwuY2EuNDU2M2MwNTEuY29uc3VsMB4XDTIyMDMwMTIyNTYxOFoX\nDTIyMDMwNDIyNTYxOFowADBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABOkztqwq\nP4SnSZ+T1JIakPeSrgcL+k30wu7rAE+xVN5lsY+iK6DAIVmHapLkOsuElI13arJa\nDaaqqdWJUG2LtqGjgfkwgfYwDgYDVR0PAQH/BAQDAgO4MB0GA1UdJQQWMBQGCCsG\nAQUFBwMCBggrBgEFBQcDATAMBgNVHRMBAf8EAjAAMCkGA1UdDgQiBCB7wCCVHVTd\nv6C07SflIf2lX1pvC1wlQIQi2zrhxaBg7TArBgNVHSMEJDAigCAMfL0aTpEwCQMh\nrD6OZMrC7lJyKSB339GwDGyU4OV3vzBfBgNVHREBAf8EVTBThlFzcGlmZmU6Ly80\nNTYzYzA1MS0wYTkyLWUwMjEtZGE2OS0wYmU5YWNkNjUxZGEuY29uc3VsL25zL2Rl\nZmF1bHQvZGMvZGMxL3N2Yy9jbGllbnQwCgYIKoZIzj0EAwIDSAAwRQIhAKKrhL0B\ny4PR/8a30JC7BmBmNWxrPSRIBaLsdhMJ9CDPAiAA7RJqkh1sc6XLx65P9FYSqDxT\nViilKSWGfQ23Ik8i1Q==\n-----END CERTIFICATE-----\n"
              },
              "private_key": {
               "inline_string": "[redacted]"
              }
             }
            ],
            "validation_context": {
             "trusted_ca": {
              "inline_string": "-----BEGIN CERTIFICATE-----\nMIICDjCCAbOgAwIBAgIBBzAKBggqhkjOPQQDAjAwMS4wLAYDVQQDEyVwcmktYWlz\nY3dnMS5jb25zdWwuY2EuNDU2M2MwNTEuY29uc3VsMB4XDTIyMDMwMTIyNTY1OFoX\nDTMyMDIyNzIyNTY1OFowMDEuMCwGA1UEAxMlcHJpLWFpc2N3ZzEuY29uc3VsLmNh\nLjQ1NjNjMDUxLmNvbnN1bDBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABDzGo4Ao\nas3SpBZl+0/WG8MWcuMgcu/VihHmxs+kRlepVWC+H9KA0IwtUTKgVtCKf7qp7dbJ\nG54R4tgv5qG6X/Cjgb0wgbowDgYDVR0PAQH/BAQDAgGGMA8GA1UdEwEB/wQFMAMB\nAf8wKQYDVR0OBCIEIAx8vRpOkTAJAyGsPo5kysLuUnIpIHff0bAMbJTg5Xe/MCsG\nA1UdIwQkMCKAIAx8vRpOkTAJAyGsPo5kysLuUnIpIHff0bAMbJTg5Xe/MD8GA1Ud\nEQQ4MDaGNHNwaWZmZTovLzQ1NjNjMDUxLTBhOTItZTAyMS1kYTY5LTBiZTlhY2Q2\nNTFkYS5jb25zdWwwCgYIKoZIzj0EAwIDSQAwRgIhALkoAuOMTRSTMQnByTCN11Uk\nOsd9eet4efD8tJtXwXppAiEAl6Fd/cWtGLT3ciEOpgxJMIEBwTKtd9xO/KJa67Cr\nZJM=\n-----END CERTIFICATE-----\n"
             }
            }
           },
           "require_client_certificate": true
          }
         }
        }
       ],
       "traffic_direction": "INBOUND"
      },
      "last_updated": "2022-03-01T22:57:32.968Z"
     }
    },
    {
     "name": "outbound_listener:127.0.0.1:15001",
     "active_state": {
      "version_info": "98862652ea5b8bd2b211d75750eda07fbfc5907deab4e4ce7b2a0619ca52485b",
      "listener": {
       "@type": "type.googleapis.com/envoy.config.listener.v3.Listener",
       "name": "outbound_listener:127.0.0.1:15001",
       "address": {
        "socket_address": {
         "address": "127.0.0.1",
         "port_value": 15001
        }
       },
       "filter_chains": [
        {
         "filter_chain_match": {
          "prefix_ranges": [
           {
            "address_prefix": "10.96.46.179",
            "prefix_len": 32
           },
           {
            "address_prefix": "240.0.0.1",
            "prefix_len": 32
           }
          ]
         },
         "filters": [
          {
           "name": "envoy.filters.network.tcp_proxy",
           "typed_config": {
            "@type": "type.googleapis.com/envoy.extensions.filters.network.tcp_proxy.v3.TcpProxy",
            "stat_prefix": "upstream.server.default.default.dc1",
            "cluster": "server.default.dc1.internal.4563c051-0a92-e021-da69-0be9acd651da.consul"
           }
          }
         ]
        },
        {
         "filters": [
          {
           "name": "envoy.filters.network.tcp_proxy",
           "typed_config": {
            "@type": "type.googleapis.com/envoy.extensions.filters.network.tcp_proxy.v3.TcpProxy",
            "stat_prefix": "upstream.original-destination",
            "cluster": "original-destination"
           }
          }
         ]
        }
       ],
       "listener_filters": [
        {
         "name": "envoy.filters.listener.original_dst"
        }
       ],
       "traffic_direction": "OUTBOUND"
      },
      "last_updated": "2022-03-01T22:57:32.971Z"
     }
    }
   ]
  },
  {
   "@type": "type.googleapis.com/envoy.admin.v3.SecretsConfigDump"
  }
 ]
}`
