// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package route53

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// A request to get a list of geographic locations that Amazon Route 53 supports
// for geolocation resource record sets.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/route53-2013-04-01/ListGeoLocationsRequest
type ListGeoLocationsInput struct {
	_ struct{} `type:"structure"`

	// (Optional) The maximum number of geolocations to be included in the response
	// body for this request. If more than maxitems geolocations remain to be listed,
	// then the value of the IsTruncated element in the response is true.
	MaxItems *string `location:"querystring" locationName:"maxitems" type:"string"`

	// The code for the continent with which you want to start listing locations
	// that Amazon Route 53 supports for geolocation. If Route 53 has already returned
	// a page or more of results, if IsTruncated is true, and if NextContinentCode
	// from the previous response has a value, enter that value in startcontinentcode
	// to return the next page of results.
	//
	// Include startcontinentcode only if you want to list continents. Don't include
	// startcontinentcode when you're listing countries or countries with their
	// subdivisions.
	StartContinentCode *string `location:"querystring" locationName:"startcontinentcode" min:"2" type:"string"`

	// The code for the country with which you want to start listing locations that
	// Amazon Route 53 supports for geolocation. If Route 53 has already returned
	// a page or more of results, if IsTruncated is true, and if NextCountryCode
	// from the previous response has a value, enter that value in startcountrycode
	// to return the next page of results.
	//
	// Route 53 uses the two-letter country codes that are specified in ISO standard
	// 3166-1 alpha-2 (https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2).
	StartCountryCode *string `location:"querystring" locationName:"startcountrycode" min:"1" type:"string"`

	// The code for the subdivision (for example, state or province) with which
	// you want to start listing locations that Amazon Route 53 supports for geolocation.
	// If Route 53 has already returned a page or more of results, if IsTruncated
	// is true, and if NextSubdivisionCode from the previous response has a value,
	// enter that value in startsubdivisioncode to return the next page of results.
	//
	// To list subdivisions of a country, you must include both startcountrycode
	// and startsubdivisioncode.
	StartSubdivisionCode *string `location:"querystring" locationName:"startsubdivisioncode" min:"1" type:"string"`
}

// String returns the string representation
func (s ListGeoLocationsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListGeoLocationsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListGeoLocationsInput"}
	if s.StartContinentCode != nil && len(*s.StartContinentCode) < 2 {
		invalidParams.Add(aws.NewErrParamMinLen("StartContinentCode", 2))
	}
	if s.StartCountryCode != nil && len(*s.StartCountryCode) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("StartCountryCode", 1))
	}
	if s.StartSubdivisionCode != nil && len(*s.StartSubdivisionCode) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("StartSubdivisionCode", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListGeoLocationsInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.MaxItems != nil {
		v := *s.MaxItems

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "maxitems", protocol.StringValue(v), metadata)
	}
	if s.StartContinentCode != nil {
		v := *s.StartContinentCode

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "startcontinentcode", protocol.StringValue(v), metadata)
	}
	if s.StartCountryCode != nil {
		v := *s.StartCountryCode

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "startcountrycode", protocol.StringValue(v), metadata)
	}
	if s.StartSubdivisionCode != nil {
		v := *s.StartSubdivisionCode

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "startsubdivisioncode", protocol.StringValue(v), metadata)
	}
	return nil
}

// A complex type containing the response information for the request.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/route53-2013-04-01/ListGeoLocationsResponse
type ListGeoLocationsOutput struct {
	_ struct{} `type:"structure"`

	// A complex type that contains one GeoLocationDetails element for each location
	// that Amazon Route 53 supports for geolocation.
	//
	// GeoLocationDetailsList is a required field
	GeoLocationDetailsList []GeoLocationDetails `locationNameList:"GeoLocationDetails" type:"list" required:"true"`

	// A value that indicates whether more locations remain to be listed after the
	// last location in this response. If so, the value of IsTruncated is true.
	// To get more values, submit another request and include the values of NextContinentCode,
	// NextCountryCode, and NextSubdivisionCode in the startcontinentcode, startcountrycode,
	// and startsubdivisioncode, as applicable.
	//
	// IsTruncated is a required field
	IsTruncated *bool `type:"boolean" required:"true"`

	// The value that you specified for MaxItems in the request.
	//
	// MaxItems is a required field
	MaxItems *string `type:"string" required:"true"`

	// If IsTruncated is true, you can make a follow-up request to display more
	// locations. Enter the value of NextContinentCode in the startcontinentcode
	// parameter in another ListGeoLocations request.
	NextContinentCode *string `min:"2" type:"string"`

	// If IsTruncated is true, you can make a follow-up request to display more
	// locations. Enter the value of NextCountryCode in the startcountrycode parameter
	// in another ListGeoLocations request.
	NextCountryCode *string `min:"1" type:"string"`

	// If IsTruncated is true, you can make a follow-up request to display more
	// locations. Enter the value of NextSubdivisionCode in the startsubdivisioncode
	// parameter in another ListGeoLocations request.
	NextSubdivisionCode *string `min:"1" type:"string"`
}

// String returns the string representation
func (s ListGeoLocationsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListGeoLocationsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.GeoLocationDetailsList != nil {
		v := s.GeoLocationDetailsList

		metadata := protocol.Metadata{ListLocationName: "GeoLocationDetails"}
		ls0 := e.List(protocol.BodyTarget, "GeoLocationDetailsList", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.IsTruncated != nil {
		v := *s.IsTruncated

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "IsTruncated", protocol.BoolValue(v), metadata)
	}
	if s.MaxItems != nil {
		v := *s.MaxItems

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "MaxItems", protocol.StringValue(v), metadata)
	}
	if s.NextContinentCode != nil {
		v := *s.NextContinentCode

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "NextContinentCode", protocol.StringValue(v), metadata)
	}
	if s.NextCountryCode != nil {
		v := *s.NextCountryCode

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "NextCountryCode", protocol.StringValue(v), metadata)
	}
	if s.NextSubdivisionCode != nil {
		v := *s.NextSubdivisionCode

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "NextSubdivisionCode", protocol.StringValue(v), metadata)
	}
	return nil
}

const opListGeoLocations = "ListGeoLocations"

// ListGeoLocationsRequest returns a request value for making API operation for
// Amazon Route 53.
//
// Retrieves a list of supported geographic locations.
//
// Countries are listed first, and continents are listed last. If Amazon Route
// 53 supports subdivisions for a country (for example, states or provinces),
// the subdivisions for that country are listed in alphabetical order immediately
// after the corresponding country.
//
//    // Example sending a request using ListGeoLocationsRequest.
//    req := client.ListGeoLocationsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/route53-2013-04-01/ListGeoLocations
func (c *Client) ListGeoLocationsRequest(input *ListGeoLocationsInput) ListGeoLocationsRequest {
	op := &aws.Operation{
		Name:       opListGeoLocations,
		HTTPMethod: "GET",
		HTTPPath:   "/2013-04-01/geolocations",
	}

	if input == nil {
		input = &ListGeoLocationsInput{}
	}

	req := c.newRequest(op, input, &ListGeoLocationsOutput{})
	return ListGeoLocationsRequest{Request: req, Input: input, Copy: c.ListGeoLocationsRequest}
}

// ListGeoLocationsRequest is the request type for the
// ListGeoLocations API operation.
type ListGeoLocationsRequest struct {
	*aws.Request
	Input *ListGeoLocationsInput
	Copy  func(*ListGeoLocationsInput) ListGeoLocationsRequest
}

// Send marshals and sends the ListGeoLocations API request.
func (r ListGeoLocationsRequest) Send(ctx context.Context) (*ListGeoLocationsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListGeoLocationsResponse{
		ListGeoLocationsOutput: r.Request.Data.(*ListGeoLocationsOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListGeoLocationsResponse is the response type for the
// ListGeoLocations API operation.
type ListGeoLocationsResponse struct {
	*ListGeoLocationsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListGeoLocations request.
func (r *ListGeoLocationsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
