// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elasticloadbalancing

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

var _ aws.Config
var _ = awsutil.Prettify

// Information about the AccessLog attribute.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticloadbalancing-2012-06-01/AccessLog
type AccessLog struct {
	_ struct{} `type:"structure"`

	// The interval for publishing the access logs. You can specify an interval
	// of either 5 minutes or 60 minutes.
	//
	// Default: 60 minutes
	EmitInterval *int64 `json:"elasticloadbalancing:AccessLog:EmitInterval" type:"integer"`

	// Specifies whether access logs are enabled for the load balancer.
	//
	// Enabled is a required field
	Enabled *bool `json:"elasticloadbalancing:AccessLog:Enabled" type:"boolean" required:"true"`

	// The name of the Amazon S3 bucket where the access logs are stored.
	S3BucketName *string `json:"elasticloadbalancing:AccessLog:S3BucketName" type:"string"`

	// The logical hierarchy you created for your Amazon S3 bucket, for example
	// my-bucket-prefix/prod. If the prefix is not provided, the log is placed at
	// the root level of the bucket.
	S3BucketPrefix *string `json:"elasticloadbalancing:AccessLog:S3BucketPrefix" type:"string"`
}

// String returns the string representation
func (s AccessLog) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AccessLog) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AccessLog"}

	if s.Enabled == nil {
		invalidParams.Add(aws.NewErrParamRequired("Enabled"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// This data type is reserved.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticloadbalancing-2012-06-01/AdditionalAttribute
type AdditionalAttribute struct {
	_ struct{} `type:"structure"`

	// This parameter is reserved.
	Key *string `json:"elasticloadbalancing:AdditionalAttribute:Key" type:"string"`

	// This parameter is reserved.
	Value *string `json:"elasticloadbalancing:AdditionalAttribute:Value" type:"string"`
}

// String returns the string representation
func (s AdditionalAttribute) String() string {
	return awsutil.Prettify(s)
}

// Information about a policy for application-controlled session stickiness.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticloadbalancing-2012-06-01/AppCookieStickinessPolicy
type AppCookieStickinessPolicy struct {
	_ struct{} `type:"structure"`

	// The name of the application cookie used for stickiness.
	CookieName *string `json:"elasticloadbalancing:AppCookieStickinessPolicy:CookieName" type:"string"`

	// The mnemonic name for the policy being created. The name must be unique within
	// a set of policies for this load balancer.
	PolicyName *string `json:"elasticloadbalancing:AppCookieStickinessPolicy:PolicyName" type:"string"`
}

// String returns the string representation
func (s AppCookieStickinessPolicy) String() string {
	return awsutil.Prettify(s)
}

// Information about the configuration of an EC2 instance.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticloadbalancing-2012-06-01/BackendServerDescription
type BackendServerDescription struct {
	_ struct{} `type:"structure"`

	// The port on which the EC2 instance is listening.
	InstancePort *int64 `json:"elasticloadbalancing:BackendServerDescription:InstancePort" min:"1" type:"integer"`

	// The names of the policies enabled for the EC2 instance.
	PolicyNames []string `json:"elasticloadbalancing:BackendServerDescription:PolicyNames" type:"list"`
}

// String returns the string representation
func (s BackendServerDescription) String() string {
	return awsutil.Prettify(s)
}

// Information about the ConnectionDraining attribute.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticloadbalancing-2012-06-01/ConnectionDraining
type ConnectionDraining struct {
	_ struct{} `type:"structure"`

	// Specifies whether connection draining is enabled for the load balancer.
	//
	// Enabled is a required field
	Enabled *bool `json:"elasticloadbalancing:ConnectionDraining:Enabled" type:"boolean" required:"true"`

	// The maximum time, in seconds, to keep the existing connections open before
	// deregistering the instances.
	Timeout *int64 `json:"elasticloadbalancing:ConnectionDraining:Timeout" type:"integer"`
}

// String returns the string representation
func (s ConnectionDraining) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ConnectionDraining) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ConnectionDraining"}

	if s.Enabled == nil {
		invalidParams.Add(aws.NewErrParamRequired("Enabled"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Information about the ConnectionSettings attribute.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticloadbalancing-2012-06-01/ConnectionSettings
type ConnectionSettings struct {
	_ struct{} `type:"structure"`

	// The time, in seconds, that the connection is allowed to be idle (no data
	// has been sent over the connection) before it is closed by the load balancer.
	//
	// IdleTimeout is a required field
	IdleTimeout *int64 `json:"elasticloadbalancing:ConnectionSettings:IdleTimeout" min:"1" type:"integer" required:"true"`
}

// String returns the string representation
func (s ConnectionSettings) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ConnectionSettings) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ConnectionSettings"}

	if s.IdleTimeout == nil {
		invalidParams.Add(aws.NewErrParamRequired("IdleTimeout"))
	}
	if s.IdleTimeout != nil && *s.IdleTimeout < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("IdleTimeout", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Information about the CrossZoneLoadBalancing attribute.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticloadbalancing-2012-06-01/CrossZoneLoadBalancing
type CrossZoneLoadBalancing struct {
	_ struct{} `type:"structure"`

	// Specifies whether cross-zone load balancing is enabled for the load balancer.
	//
	// Enabled is a required field
	Enabled *bool `json:"elasticloadbalancing:CrossZoneLoadBalancing:Enabled" type:"boolean" required:"true"`
}

// String returns the string representation
func (s CrossZoneLoadBalancing) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CrossZoneLoadBalancing) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CrossZoneLoadBalancing"}

	if s.Enabled == nil {
		invalidParams.Add(aws.NewErrParamRequired("Enabled"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Information about a health check.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticloadbalancing-2012-06-01/HealthCheck
type HealthCheck struct {
	_ struct{} `type:"structure"`

	// The number of consecutive health checks successes required before moving
	// the instance to the Healthy state.
	//
	// HealthyThreshold is a required field
	HealthyThreshold *int64 `json:"elasticloadbalancing:HealthCheck:HealthyThreshold" min:"2" type:"integer" required:"true"`

	// The approximate interval, in seconds, between health checks of an individual
	// instance.
	//
	// Interval is a required field
	Interval *int64 `json:"elasticloadbalancing:HealthCheck:Interval" min:"5" type:"integer" required:"true"`

	// The instance being checked. The protocol is either TCP, HTTP, HTTPS, or SSL.
	// The range of valid ports is one (1) through 65535.
	//
	// TCP is the default, specified as a TCP: port pair, for example "TCP:5000".
	// In this case, a health check simply attempts to open a TCP connection to
	// the instance on the specified port. Failure to connect within the configured
	// timeout is considered unhealthy.
	//
	// SSL is also specified as SSL: port pair, for example, SSL:5000.
	//
	// For HTTP/HTTPS, you must include a ping path in the string. HTTP is specified
	// as a HTTP:port;/;PathToPing; grouping, for example "HTTP:80/weather/us/wa/seattle".
	// In this case, a HTTP GET request is issued to the instance on the given port
	// and path. Any answer other than "200 OK" within the timeout period is considered
	// unhealthy.
	//
	// The total length of the HTTP ping target must be 1024 16-bit Unicode characters
	// or less.
	//
	// Target is a required field
	Target *string `json:"elasticloadbalancing:HealthCheck:Target" type:"string" required:"true"`

	// The amount of time, in seconds, during which no response means a failed health
	// check.
	//
	// This value must be less than the Interval value.
	//
	// Timeout is a required field
	Timeout *int64 `json:"elasticloadbalancing:HealthCheck:Timeout" min:"2" type:"integer" required:"true"`

	// The number of consecutive health check failures required before moving the
	// instance to the Unhealthy state.
	//
	// UnhealthyThreshold is a required field
	UnhealthyThreshold *int64 `json:"elasticloadbalancing:HealthCheck:UnhealthyThreshold" min:"2" type:"integer" required:"true"`
}

// String returns the string representation
func (s HealthCheck) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *HealthCheck) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "HealthCheck"}

	if s.HealthyThreshold == nil {
		invalidParams.Add(aws.NewErrParamRequired("HealthyThreshold"))
	}
	if s.HealthyThreshold != nil && *s.HealthyThreshold < 2 {
		invalidParams.Add(aws.NewErrParamMinValue("HealthyThreshold", 2))
	}

	if s.Interval == nil {
		invalidParams.Add(aws.NewErrParamRequired("Interval"))
	}
	if s.Interval != nil && *s.Interval < 5 {
		invalidParams.Add(aws.NewErrParamMinValue("Interval", 5))
	}

	if s.Target == nil {
		invalidParams.Add(aws.NewErrParamRequired("Target"))
	}

	if s.Timeout == nil {
		invalidParams.Add(aws.NewErrParamRequired("Timeout"))
	}
	if s.Timeout != nil && *s.Timeout < 2 {
		invalidParams.Add(aws.NewErrParamMinValue("Timeout", 2))
	}

	if s.UnhealthyThreshold == nil {
		invalidParams.Add(aws.NewErrParamRequired("UnhealthyThreshold"))
	}
	if s.UnhealthyThreshold != nil && *s.UnhealthyThreshold < 2 {
		invalidParams.Add(aws.NewErrParamMinValue("UnhealthyThreshold", 2))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// The ID of an EC2 instance.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticloadbalancing-2012-06-01/Instance
type Instance struct {
	_ struct{} `type:"structure"`

	// The instance ID.
	InstanceId *string `json:"elasticloadbalancing:Instance:InstanceId" type:"string"`
}

// String returns the string representation
func (s Instance) String() string {
	return awsutil.Prettify(s)
}

// Information about the state of an EC2 instance.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticloadbalancing-2012-06-01/InstanceState
type InstanceState struct {
	_ struct{} `type:"structure"`

	// A description of the instance state. This string can contain one or more
	// of the following messages.
	//
	//    * N/A
	//
	//    * A transient error occurred. Please try again later.
	//
	//    * Instance has failed at least the UnhealthyThreshold number of health
	//    checks consecutively.
	//
	//    * Instance has not passed the configured HealthyThreshold number of health
	//    checks consecutively.
	//
	//    * Instance registration is still in progress.
	//
	//    * Instance is in the EC2 Availability Zone for which LoadBalancer is not
	//    configured to route traffic to.
	//
	//    * Instance is not currently registered with the LoadBalancer.
	//
	//    * Instance deregistration currently in progress.
	//
	//    * Disable Availability Zone is currently in progress.
	//
	//    * Instance is in pending state.
	//
	//    * Instance is in stopped state.
	//
	//    * Instance is in terminated state.
	Description *string `json:"elasticloadbalancing:InstanceState:Description" type:"string"`

	// The ID of the instance.
	InstanceId *string `json:"elasticloadbalancing:InstanceState:InstanceId" type:"string"`

	// Information about the cause of OutOfService instances. Specifically, whether
	// the cause is Elastic Load Balancing or the instance.
	//
	// Valid values: ELB | Instance | N/A
	ReasonCode *string `json:"elasticloadbalancing:InstanceState:ReasonCode" type:"string"`

	// The current state of the instance.
	//
	// Valid values: InService | OutOfService | Unknown
	State *string `json:"elasticloadbalancing:InstanceState:State" type:"string"`
}

// String returns the string representation
func (s InstanceState) String() string {
	return awsutil.Prettify(s)
}

// Information about a policy for duration-based session stickiness.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticloadbalancing-2012-06-01/LBCookieStickinessPolicy
type LBCookieStickinessPolicy struct {
	_ struct{} `type:"structure"`

	// The time period, in seconds, after which the cookie should be considered
	// stale. If this parameter is not specified, the stickiness session lasts for
	// the duration of the browser session.
	CookieExpirationPeriod *int64 `json:"elasticloadbalancing:LBCookieStickinessPolicy:CookieExpirationPeriod" type:"long"`

	// The name of the policy. This name must be unique within the set of policies
	// for this load balancer.
	PolicyName *string `json:"elasticloadbalancing:LBCookieStickinessPolicy:PolicyName" type:"string"`
}

// String returns the string representation
func (s LBCookieStickinessPolicy) String() string {
	return awsutil.Prettify(s)
}

// Information about an Elastic Load Balancing resource limit for your AWS account.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticloadbalancing-2012-06-01/Limit
type Limit struct {
	_ struct{} `type:"structure"`

	// The maximum value of the limit.
	Max *string `json:"elasticloadbalancing:Limit:Max" type:"string"`

	// The name of the limit. The possible values are:
	//
	//    * classic-listeners
	//
	//    * classic-load-balancers
	//
	//    * classic-registered-instances
	Name *string `json:"elasticloadbalancing:Limit:Name" type:"string"`
}

// String returns the string representation
func (s Limit) String() string {
	return awsutil.Prettify(s)
}

// Information about a listener.
//
// For information about the protocols and the ports supported by Elastic Load
// Balancing, see Listeners for Your Classic Load Balancer (http://docs.aws.amazon.com/elasticloadbalancing/latest/classic/elb-listener-config.html)
// in the Classic Load Balancers Guide.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticloadbalancing-2012-06-01/Listener
type Listener struct {
	_ struct{} `type:"structure"`

	// The port on which the instance is listening.
	//
	// InstancePort is a required field
	InstancePort *int64 `json:"elasticloadbalancing:Listener:InstancePort" min:"1" type:"integer" required:"true"`

	// The protocol to use for routing traffic to instances: HTTP, HTTPS, TCP, or
	// SSL.
	//
	// If the front-end protocol is HTTP, HTTPS, TCP, or SSL, InstanceProtocol must
	// be at the same protocol.
	//
	// If there is another listener with the same InstancePort whose InstanceProtocol
	// is secure, (HTTPS or SSL), the listener's InstanceProtocol must also be secure.
	//
	// If there is another listener with the same InstancePort whose InstanceProtocol
	// is HTTP or TCP, the listener's InstanceProtocol must be HTTP or TCP.
	InstanceProtocol *string `json:"elasticloadbalancing:Listener:InstanceProtocol" type:"string"`

	// The port on which the load balancer is listening. On EC2-VPC, you can specify
	// any port from the range 1-65535. On EC2-Classic, you can specify any port
	// from the following list: 25, 80, 443, 465, 587, 1024-65535.
	//
	// LoadBalancerPort is a required field
	LoadBalancerPort *int64 `json:"elasticloadbalancing:Listener:LoadBalancerPort" type:"integer" required:"true"`

	// The load balancer transport protocol to use for routing: HTTP, HTTPS, TCP,
	// or SSL.
	//
	// Protocol is a required field
	Protocol *string `json:"elasticloadbalancing:Listener:Protocol" type:"string" required:"true"`

	// The Amazon Resource Name (ARN) of the server certificate.
	SSLCertificateId *string `json:"elasticloadbalancing:Listener:SSLCertificateId" type:"string"`
}

// String returns the string representation
func (s Listener) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *Listener) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "Listener"}

	if s.InstancePort == nil {
		invalidParams.Add(aws.NewErrParamRequired("InstancePort"))
	}
	if s.InstancePort != nil && *s.InstancePort < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("InstancePort", 1))
	}

	if s.LoadBalancerPort == nil {
		invalidParams.Add(aws.NewErrParamRequired("LoadBalancerPort"))
	}

	if s.Protocol == nil {
		invalidParams.Add(aws.NewErrParamRequired("Protocol"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// The policies enabled for a listener.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticloadbalancing-2012-06-01/ListenerDescription
type ListenerDescription struct {
	_ struct{} `type:"structure"`

	// The listener.
	Listener *Listener `json:"elasticloadbalancing:ListenerDescription:Listener" type:"structure"`

	// The policies. If there are no policies enabled, the list is empty.
	PolicyNames []string `json:"elasticloadbalancing:ListenerDescription:PolicyNames" type:"list"`
}

// String returns the string representation
func (s ListenerDescription) String() string {
	return awsutil.Prettify(s)
}

// The attributes for a load balancer.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticloadbalancing-2012-06-01/LoadBalancerAttributes
type LoadBalancerAttributes struct {
	_ struct{} `type:"structure"`

	// If enabled, the load balancer captures detailed information of all requests
	// and delivers the information to the Amazon S3 bucket that you specify.
	//
	// For more information, see Enable Access Logs (http://docs.aws.amazon.com/elasticloadbalancing/latest/classic/enable-access-logs.html)
	// in the Classic Load Balancers Guide.
	AccessLog *AccessLog `json:"elasticloadbalancing:LoadBalancerAttributes:AccessLog" type:"structure"`

	// This parameter is reserved.
	AdditionalAttributes []AdditionalAttribute `json:"elasticloadbalancing:LoadBalancerAttributes:AdditionalAttributes" type:"list"`

	// If enabled, the load balancer allows existing requests to complete before
	// the load balancer shifts traffic away from a deregistered or unhealthy instance.
	//
	// For more information, see Configure Connection Draining (http://docs.aws.amazon.com/elasticloadbalancing/latest/classic/config-conn-drain.html)
	// in the Classic Load Balancers Guide.
	ConnectionDraining *ConnectionDraining `json:"elasticloadbalancing:LoadBalancerAttributes:ConnectionDraining" type:"structure"`

	// If enabled, the load balancer allows the connections to remain idle (no data
	// is sent over the connection) for the specified duration.
	//
	// By default, Elastic Load Balancing maintains a 60-second idle connection
	// timeout for both front-end and back-end connections of your load balancer.
	// For more information, see Configure Idle Connection Timeout (http://docs.aws.amazon.com/elasticloadbalancing/latest/classic/config-idle-timeout.html)
	// in the Classic Load Balancers Guide.
	ConnectionSettings *ConnectionSettings `json:"elasticloadbalancing:LoadBalancerAttributes:ConnectionSettings" type:"structure"`

	// If enabled, the load balancer routes the request traffic evenly across all
	// instances regardless of the Availability Zones.
	//
	// For more information, see Configure Cross-Zone Load Balancing (http://docs.aws.amazon.com/elasticloadbalancing/latest/classic/enable-disable-crosszone-lb.html)
	// in the Classic Load Balancers Guide.
	CrossZoneLoadBalancing *CrossZoneLoadBalancing `json:"elasticloadbalancing:LoadBalancerAttributes:CrossZoneLoadBalancing" type:"structure"`
}

// String returns the string representation
func (s LoadBalancerAttributes) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *LoadBalancerAttributes) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "LoadBalancerAttributes"}
	if s.AccessLog != nil {
		if err := s.AccessLog.Validate(); err != nil {
			invalidParams.AddNested("AccessLog", err.(aws.ErrInvalidParams))
		}
	}
	if s.ConnectionDraining != nil {
		if err := s.ConnectionDraining.Validate(); err != nil {
			invalidParams.AddNested("ConnectionDraining", err.(aws.ErrInvalidParams))
		}
	}
	if s.ConnectionSettings != nil {
		if err := s.ConnectionSettings.Validate(); err != nil {
			invalidParams.AddNested("ConnectionSettings", err.(aws.ErrInvalidParams))
		}
	}
	if s.CrossZoneLoadBalancing != nil {
		if err := s.CrossZoneLoadBalancing.Validate(); err != nil {
			invalidParams.AddNested("CrossZoneLoadBalancing", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Information about a load balancer.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticloadbalancing-2012-06-01/LoadBalancerDescription
type LoadBalancerDescription struct {
	_ struct{} `type:"structure"`

	// The Availability Zones for the load balancer.
	AvailabilityZones []string `json:"elasticloadbalancing:LoadBalancerDescription:AvailabilityZones" type:"list"`

	// Information about your EC2 instances.
	BackendServerDescriptions []BackendServerDescription `json:"elasticloadbalancing:LoadBalancerDescription:BackendServerDescriptions" type:"list"`

	// The DNS name of the load balancer.
	//
	// For more information, see Configure a Custom Domain Name (http://docs.aws.amazon.com/elasticloadbalancing/latest/classic/using-domain-names-with-elb.html)
	// in the Classic Load Balancers Guide.
	CanonicalHostedZoneName *string `json:"elasticloadbalancing:LoadBalancerDescription:CanonicalHostedZoneName" type:"string"`

	// The ID of the Amazon Route 53 hosted zone for the load balancer.
	CanonicalHostedZoneNameID *string `json:"elasticloadbalancing:LoadBalancerDescription:CanonicalHostedZoneNameID" type:"string"`

	// The date and time the load balancer was created.
	CreatedTime *time.Time `json:"elasticloadbalancing:LoadBalancerDescription:CreatedTime" type:"timestamp" timestampFormat:"iso8601"`

	// The DNS name of the load balancer.
	DNSName *string `json:"elasticloadbalancing:LoadBalancerDescription:DNSName" type:"string"`

	// Information about the health checks conducted on the load balancer.
	HealthCheck *HealthCheck `json:"elasticloadbalancing:LoadBalancerDescription:HealthCheck" type:"structure"`

	// The IDs of the instances for the load balancer.
	Instances []Instance `json:"elasticloadbalancing:LoadBalancerDescription:Instances" type:"list"`

	// The listeners for the load balancer.
	ListenerDescriptions []ListenerDescription `json:"elasticloadbalancing:LoadBalancerDescription:ListenerDescriptions" type:"list"`

	// The name of the load balancer.
	LoadBalancerName *string `json:"elasticloadbalancing:LoadBalancerDescription:LoadBalancerName" type:"string"`

	// The policies defined for the load balancer.
	Policies *Policies `json:"elasticloadbalancing:LoadBalancerDescription:Policies" type:"structure"`

	// The type of load balancer. Valid only for load balancers in a VPC.
	//
	// If Scheme is internet-facing, the load balancer has a public DNS name that
	// resolves to a public IP address.
	//
	// If Scheme is internal, the load balancer has a public DNS name that resolves
	// to a private IP address.
	Scheme *string `json:"elasticloadbalancing:LoadBalancerDescription:Scheme" type:"string"`

	// The security groups for the load balancer. Valid only for load balancers
	// in a VPC.
	SecurityGroups []string `json:"elasticloadbalancing:LoadBalancerDescription:SecurityGroups" type:"list"`

	// The security group for the load balancer, which you can use as part of your
	// inbound rules for your registered instances. To only allow traffic from load
	// balancers, add a security group rule that specifies this source security
	// group as the inbound source.
	SourceSecurityGroup *SourceSecurityGroup `json:"elasticloadbalancing:LoadBalancerDescription:SourceSecurityGroup" type:"structure"`

	// The IDs of the subnets for the load balancer.
	Subnets []string `json:"elasticloadbalancing:LoadBalancerDescription:Subnets" type:"list"`

	// The ID of the VPC for the load balancer.
	VPCId *string `json:"elasticloadbalancing:LoadBalancerDescription:VPCId" type:"string"`
}

// String returns the string representation
func (s LoadBalancerDescription) String() string {
	return awsutil.Prettify(s)
}

// The policies for a load balancer.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticloadbalancing-2012-06-01/Policies
type Policies struct {
	_ struct{} `type:"structure"`

	// The stickiness policies created using CreateAppCookieStickinessPolicy.
	AppCookieStickinessPolicies []AppCookieStickinessPolicy `json:"elasticloadbalancing:Policies:AppCookieStickinessPolicies" type:"list"`

	// The stickiness policies created using CreateLBCookieStickinessPolicy.
	LBCookieStickinessPolicies []LBCookieStickinessPolicy `json:"elasticloadbalancing:Policies:LBCookieStickinessPolicies" type:"list"`

	// The policies other than the stickiness policies.
	OtherPolicies []string `json:"elasticloadbalancing:Policies:OtherPolicies" type:"list"`
}

// String returns the string representation
func (s Policies) String() string {
	return awsutil.Prettify(s)
}

// Information about a policy attribute.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticloadbalancing-2012-06-01/PolicyAttribute
type PolicyAttribute struct {
	_ struct{} `type:"structure"`

	// The name of the attribute.
	AttributeName *string `json:"elasticloadbalancing:PolicyAttribute:AttributeName" type:"string"`

	// The value of the attribute.
	AttributeValue *string `json:"elasticloadbalancing:PolicyAttribute:AttributeValue" type:"string"`
}

// String returns the string representation
func (s PolicyAttribute) String() string {
	return awsutil.Prettify(s)
}

// Information about a policy attribute.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticloadbalancing-2012-06-01/PolicyAttributeDescription
type PolicyAttributeDescription struct {
	_ struct{} `type:"structure"`

	// The name of the attribute.
	AttributeName *string `json:"elasticloadbalancing:PolicyAttributeDescription:AttributeName" type:"string"`

	// The value of the attribute.
	AttributeValue *string `json:"elasticloadbalancing:PolicyAttributeDescription:AttributeValue" type:"string"`
}

// String returns the string representation
func (s PolicyAttributeDescription) String() string {
	return awsutil.Prettify(s)
}

// Information about a policy attribute type.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticloadbalancing-2012-06-01/PolicyAttributeTypeDescription
type PolicyAttributeTypeDescription struct {
	_ struct{} `type:"structure"`

	// The name of the attribute.
	AttributeName *string `json:"elasticloadbalancing:PolicyAttributeTypeDescription:AttributeName" type:"string"`

	// The type of the attribute. For example, Boolean or Integer.
	AttributeType *string `json:"elasticloadbalancing:PolicyAttributeTypeDescription:AttributeType" type:"string"`

	// The cardinality of the attribute.
	//
	// Valid values:
	//
	//    * ONE(1) : Single value required
	//
	//    * ZERO_OR_ONE(0..1) : Up to one value is allowed
	//
	//    * ZERO_OR_MORE(0..*) : Optional. Multiple values are allowed
	//
	//    * ONE_OR_MORE(1..*0) : Required. Multiple values are allowed
	Cardinality *string `json:"elasticloadbalancing:PolicyAttributeTypeDescription:Cardinality" type:"string"`

	// The default value of the attribute, if applicable.
	DefaultValue *string `json:"elasticloadbalancing:PolicyAttributeTypeDescription:DefaultValue" type:"string"`

	// A description of the attribute.
	Description *string `json:"elasticloadbalancing:PolicyAttributeTypeDescription:Description" type:"string"`
}

// String returns the string representation
func (s PolicyAttributeTypeDescription) String() string {
	return awsutil.Prettify(s)
}

// Information about a policy.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticloadbalancing-2012-06-01/PolicyDescription
type PolicyDescription struct {
	_ struct{} `type:"structure"`

	// The policy attributes.
	PolicyAttributeDescriptions []PolicyAttributeDescription `json:"elasticloadbalancing:PolicyDescription:PolicyAttributeDescriptions" type:"list"`

	// The name of the policy.
	PolicyName *string `json:"elasticloadbalancing:PolicyDescription:PolicyName" type:"string"`

	// The name of the policy type.
	PolicyTypeName *string `json:"elasticloadbalancing:PolicyDescription:PolicyTypeName" type:"string"`
}

// String returns the string representation
func (s PolicyDescription) String() string {
	return awsutil.Prettify(s)
}

// Information about a policy type.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticloadbalancing-2012-06-01/PolicyTypeDescription
type PolicyTypeDescription struct {
	_ struct{} `type:"structure"`

	// A description of the policy type.
	Description *string `json:"elasticloadbalancing:PolicyTypeDescription:Description" type:"string"`

	// The description of the policy attributes associated with the policies defined
	// by Elastic Load Balancing.
	PolicyAttributeTypeDescriptions []PolicyAttributeTypeDescription `json:"elasticloadbalancing:PolicyTypeDescription:PolicyAttributeTypeDescriptions" type:"list"`

	// The name of the policy type.
	PolicyTypeName *string `json:"elasticloadbalancing:PolicyTypeDescription:PolicyTypeName" type:"string"`
}

// String returns the string representation
func (s PolicyTypeDescription) String() string {
	return awsutil.Prettify(s)
}

// Information about a source security group.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticloadbalancing-2012-06-01/SourceSecurityGroup
type SourceSecurityGroup struct {
	_ struct{} `type:"structure"`

	// The name of the security group.
	GroupName *string `json:"elasticloadbalancing:SourceSecurityGroup:GroupName" type:"string"`

	// The owner of the security group.
	OwnerAlias *string `json:"elasticloadbalancing:SourceSecurityGroup:OwnerAlias" type:"string"`
}

// String returns the string representation
func (s SourceSecurityGroup) String() string {
	return awsutil.Prettify(s)
}

// Information about a tag.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticloadbalancing-2012-06-01/Tag
type Tag struct {
	_ struct{} `type:"structure"`

	// The key of the tag.
	//
	// Key is a required field
	Key *string `json:"elasticloadbalancing:Tag:Key" min:"1" type:"string" required:"true"`

	// The value of the tag.
	Value *string `json:"elasticloadbalancing:Tag:Value" type:"string"`
}

// String returns the string representation
func (s Tag) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *Tag) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "Tag"}

	if s.Key == nil {
		invalidParams.Add(aws.NewErrParamRequired("Key"))
	}
	if s.Key != nil && len(*s.Key) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Key", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// The tags associated with a load balancer.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticloadbalancing-2012-06-01/TagDescription
type TagDescription struct {
	_ struct{} `type:"structure"`

	// The name of the load balancer.
	LoadBalancerName *string `json:"elasticloadbalancing:TagDescription:LoadBalancerName" type:"string"`

	// The tags.
	Tags []Tag `json:"elasticloadbalancing:TagDescription:Tags" min:"1" type:"list"`
}

// String returns the string representation
func (s TagDescription) String() string {
	return awsutil.Prettify(s)
}

// The key of a tag.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticloadbalancing-2012-06-01/TagKeyOnly
type TagKeyOnly struct {
	_ struct{} `type:"structure"`

	// The name of the key.
	Key *string `json:"elasticloadbalancing:TagKeyOnly:Key" min:"1" type:"string"`
}

// String returns the string representation
func (s TagKeyOnly) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *TagKeyOnly) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "TagKeyOnly"}
	if s.Key != nil && len(*s.Key) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Key", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}
