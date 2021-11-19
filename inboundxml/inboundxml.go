// Copyright 2017 Avaya Cloud Inc. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package inboundxml

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"unsafe"

	"github.com/jbussdieker/golibxml"
	"github.com/krolaw/xsd"
)

var ZANG_SCHEMA = `<?xml version="1.0"?>
<xsd:schema xmlns:xsd="http://www.w3.org/2001/XMLSchema">
    <xsd:simpleType name="reject_modes">
        <xsd:restriction base="xsd:string">
            <xsd:enumeration value="busy"/>
            <xsd:enumeration value="rejected"/>
        </xsd:restriction>
    </xsd:simpleType>
    <xsd:simpleType name="say_voice">
        <xsd:restriction base="xsd:string">
            <xsd:enumeration value="female"/>
            <xsd:enumeration value="male"/>
        </xsd:restriction>
    </xsd:simpleType>
    <xsd:simpleType name="http_methods">
        <xsd:restriction base="xsd:string">
            <xsd:enumeration value="GET"/>
            <xsd:enumeration value="POST"/>
            <xsd:enumeration value="get"/>
            <xsd:enumeration value="post"/>
        </xsd:restriction>
    </xsd:simpleType>
    <xsd:simpleType name="booleans">
        <xsd:restriction base="xsd:string">
            <xsd:enumeration value="True"/>
            <xsd:enumeration value="true"/>
            <xsd:enumeration value="TRUE"/>
            <xsd:enumeration value="1"/>
            <xsd:enumeration value="False"/>
            <xsd:enumeration value="false"/>
            <xsd:enumeration value="FALSE"/>
            <xsd:enumeration value="0"/>
        </xsd:restriction>
    </xsd:simpleType>
    <xsd:simpleType name="positive_integer">
        <xsd:restriction base="xsd:integer">
            <xsd:minExclusive value="0"/>
        </xsd:restriction>
    </xsd:simpleType>
    <xsd:simpleType name="boolean_or_positive_integer">
        <xsd:union memberTypes="positive_integer booleans"/>
    </xsd:simpleType>
    <xsd:simpleType name="zangDirection">
        <xsd:restriction base="xsd:string">
            <xsd:enumeration value=""/>
            <!-- Backwards compatible with bothLegs -->
            <xsd:enumeration value="in"/>
            <xsd:enumeration value="In"/>
            <xsd:enumeration value="IN"/>
            <xsd:enumeration value="out"/>
            <xsd:enumeration value="Out"/>
            <xsd:enumeration value="OUT"/>
            <xsd:enumeration value="both"/>
            <xsd:enumeration value="Both"/>
            <xsd:enumeration value="BOTH"/>
            <xsd:enumeration value="inlow"/>
        </xsd:restriction>
    </xsd:simpleType>
    <xsd:simpleType name="sample_rate">
        <xsd:restriction base="xsd:integer">
            <xsd:enumeration value="8000"/>
            <xsd:enumeration value="16000"/>
            <xsd:enumeration value="32000"/>
            <xsd:enumeration value="48000"/>
        </xsd:restriction>
    </xsd:simpleType>
    <xsd:simpleType name="tts_languages">
        <xsd:restriction base="xsd:string">
            <xsd:enumeration value="en"/>
            <xsd:enumeration value="en-gb"/>
            <xsd:enumeration value="es"/>
            <xsd:enumeration value="fr"/>
            <xsd:enumeration value="it"/>
            <xsd:enumeration value="de"/>
        </xsd:restriction>
    </xsd:simpleType>
    <xsd:simpleType name="bcp_languages">
        <xsd:restriction base="xsd:string">            
            <xsd:enumeration value="af-ZA"/>           
            <xsd:enumeration value="am-ET"/>           
            <xsd:enumeration value="hy-AM"/>           
            <xsd:enumeration value="az-AZ"/>           
            <xsd:enumeration value="id-ID"/>           
            <xsd:enumeration value="ms-MY"/>           
            <xsd:enumeration value="bn-BD"/>           
            <xsd:enumeration value="bn-IN"/>           
            <xsd:enumeration value="ca-ES"/>           
            <xsd:enumeration value="cs-CZ"/>           
            <xsd:enumeration value="da-DK"/>           
            <xsd:enumeration value="de-DE"/>           
            <xsd:enumeration value="en-AU"/>           
            <xsd:enumeration value="en-CA"/>
            <xsd:enumeration value="en-GH"/>
            <xsd:enumeration value="en-GB"/>
            <xsd:enumeration value="en-IN"/>
            <xsd:enumeration value="en-IE"/>
            <xsd:enumeration value="en-KE"/>
            <xsd:enumeration value="en-NZ"/>
            <xsd:enumeration value="en-NG"/>
            <xsd:enumeration value="en-PH"/>
            <xsd:enumeration value="en-ZA"/>
            <xsd:enumeration value="en-TZ"/>
            <xsd:enumeration value="en-US"/>
            <xsd:enumeration value="es-AR"/>
            <xsd:enumeration value="es-BO"/>
            <xsd:enumeration value="es-CL"/>
            <xsd:enumeration value="es-CO"/>
            <xsd:enumeration value="es-CR"/>
            <xsd:enumeration value="es-EC"/>
            <xsd:enumeration value="es-SV"/>
            <xsd:enumeration value="es-ES"/>
            <xsd:enumeration value="es-US"/>
            <xsd:enumeration value="es-GT"/>
            <xsd:enumeration value="es-HN"/>
            <xsd:enumeration value="es-MX"/>
            <xsd:enumeration value="es-NI"/>
            <xsd:enumeration value="es-PA"/>
            <xsd:enumeration value="es-PY"/>
            <xsd:enumeration value="es-PE"/>
            <xsd:enumeration value="es-PR"/>
            <xsd:enumeration value="es-DO"/>
            <xsd:enumeration value="es-UY"/>
            <xsd:enumeration value="es-VE"/>
            <xsd:enumeration value="eu-ES"/>
            <xsd:enumeration value="il-PH"/>
            <xsd:enumeration value="fr-CA"/>
            <xsd:enumeration value="fr-FR"/>
            <xsd:enumeration value="gl-ES"/>
            <xsd:enumeration value="ka-GE"/>
            <xsd:enumeration value="gu-IN"/>
            <xsd:enumeration value="hr-HR"/>
            <xsd:enumeration value="zu-ZA"/>
            <xsd:enumeration value="is-IS"/>
            <xsd:enumeration value="it-IT"/>
            <xsd:enumeration value="jv-ID"/>
            <xsd:enumeration value="kn-IN"/>
            <xsd:enumeration value="km-KH"/>
            <xsd:enumeration value="lo-LA"/>
            <xsd:enumeration value="lv-LV"/>
            <xsd:enumeration value="lt-LT"/>
            <xsd:enumeration value="hu-HU"/>
            <xsd:enumeration value="ml-IN"/>
            <xsd:enumeration value="mr-IN"/>
            <xsd:enumeration value="nl-NL"/>
            <xsd:enumeration value="ne-NP"/>
            <xsd:enumeration value="nb-NO"/>
            <xsd:enumeration value="pl-PL"/>
            <xsd:enumeration value="pt-BR"/>
            <xsd:enumeration value="pt-PT"/>
            <xsd:enumeration value="ro-RO"/>
            <xsd:enumeration value="si-LK"/>
            <xsd:enumeration value="sk-SK"/>
            <xsd:enumeration value="sl-SI"/>
            <xsd:enumeration value="su-ID"/>
            <xsd:enumeration value="sw-TZ"/>
            <xsd:enumeration value="sw-KE"/>
            <xsd:enumeration value="fi-FI"/>
            <xsd:enumeration value="sv-SE"/>
            <xsd:enumeration value="ta-IN"/>
            <xsd:enumeration value="ta-SG"/>
            <xsd:enumeration value="ta-LK"/>
            <xsd:enumeration value="ta-MY"/>
            <xsd:enumeration value="te-IN"/>
            <xsd:enumeration value="vi-VN"/>
            <xsd:enumeration value="tr-TR"/>
            <xsd:enumeration value="ur-PK"/>
            <xsd:enumeration value="ur-IN"/>
            <xsd:enumeration value="el-GR"/>
            <xsd:enumeration value="bg-BG"/>
            <xsd:enumeration value="ru-RU"/>
            <xsd:enumeration value="sr-RS"/>
            <xsd:enumeration value="uk-UA"/>
            <xsd:enumeration value="he-IL"/>
            <xsd:enumeration value="ar-IL"/>
            <xsd:enumeration value="ar-JO"/>
            <xsd:enumeration value="ar-AE"/>
            <xsd:enumeration value="ar-BH"/>
            <xsd:enumeration value="ar-DZ"/>
            <xsd:enumeration value="ar-SA"/>
            <xsd:enumeration value="ar-IQ"/>
            <xsd:enumeration value="ar-KW"/>
            <xsd:enumeration value="ar-MA"/>
            <xsd:enumeration value="ar-TN"/>
            <xsd:enumeration value="ar-OM"/>
            <xsd:enumeration value="ar-PS"/>
            <xsd:enumeration value="ar-QA"/>
            <xsd:enumeration value="ar-LB"/>
            <xsd:enumeration value="ar-EG"/>
            <xsd:enumeration value="fa-IR"/>
            <xsd:enumeration value="hi-IN"/>
            <xsd:enumeration value="th-TH"/>
            <xsd:enumeration value="ko-KR"/>
            <xsd:enumeration value="cmn-Hant-TW"/>
            <xsd:enumeration value="yue-Hant-HK"/>
            <xsd:enumeration value="ja-JP"/>
            <xsd:enumeration value="cmn-Hans-HK"/>
            <xsd:enumeration value="cmn-Hans-CN"/>
        </xsd:restriction>
    </xsd:simpleType>
    <xsd:simpleType name="transcription_quality">
        <xsd:restriction base="xsd:string">
            <xsd:enumeration value="auto"/>
            <xsd:enumeration value="silver"/>
            <xsd:enumeration value="gold"/>
            <xsd:enumeration value="platinum"/>
            <xsd:enumeration value="manual"/>
            <xsd:enumeration value="AUTO"/>
            <xsd:enumeration value="SILVER"/>
            <xsd:enumeration value="GOLD"/>
            <xsd:enumeration value="PLATINUM"/>
            <xsd:enumeration value="MANUAL"/>
            <xsd:enumeration value="hybrid"/>
            <xsd:enumeration value="Hybrid"/>
            <xsd:enumeration value="HYBRID"/>
            <xsd:enumeration value="keywords"/>
            <xsd:enumeration value="Keywords"/>
            <xsd:enumeration value="KEYWORDS"/>
            <xsd:enumeration value="subtitles"/>
            <xsd:enumeration value="Subtitles"/>
            <xsd:enumeration value="SUBTITLES"/>
        </xsd:restriction>
    </xsd:simpleType>
    <xsd:simpleType name="if_machine">
        <xsd:restriction base="xsd:string">
            <xsd:enumeration value="continue"/>
            <xsd:enumeration value="Continue"/>
            <xsd:enumeration value="CONTINUE"/>
            <xsd:enumeration value="hangup"/>
            <xsd:enumeration value="Hangup"/>
            <xsd:enumeration value="HANGUP"/>
            <xsd:enumeration value="HangUp"/>
            <xsd:enumeration value="Redirect"/>
            <xsd:enumeration value="redirect"/>
            <xsd:enumeration value="REDIRECT"/>
        </xsd:restriction>
    </xsd:simpleType>
    <xsd:simpleType id="anyURINonBlank" name="anyURINonBlank">
        <xsd:restriction base="xsd:anyURI">
            <xsd:minLength value="1"/>
        </xsd:restriction>
    </xsd:simpleType>
    <xsd:simpleType name="gather_input">
        <xsd:restriction base="xsd:string">
            <xsd:enumeration value="dtmf"/>
            <xsd:enumeration value="speech"/>
            <xsd:enumeration value="speech dtmf"/>
        </xsd:restriction>
    </xsd:simpleType>

    <xsd:element name="Conference">
        <xsd:complexType mixed="true">
            <xsd:attribute name="muted" type="booleans" use="optional"/>
            <xsd:attribute name="beep" type="booleans" use="optional"/>
            <xsd:attribute name="startConferenceOnEnter" type="booleans" use="optional"/>
            <xsd:attribute name="endConferenceOnExit" type="booleans" use="optional"/>
            <xsd:attribute name="maxParticipants" type="xsd:integer" use="optional"/>
            <xsd:attribute name="timeLimit" type="xsd:integer" use="optional"/>
            <xsd:attribute name="waitUrl" type="xsd:anyURI" use="optional"/>
            <xsd:attribute name="waitMethod" type="http_methods" use="optional"/>
            <xsd:attribute name="hangupOnStar" type="xsd:string" use="optional"/>
            <xsd:attribute name="callbackUrl" type="xsd:anyURI" use="optional"/>
            <xsd:attribute name="callbackMethod" type="http_methods" use="optional"/>
            <xsd:attribute name="method" type="http_methods" use="optional"/>
            <xsd:attribute name="waitSound" type="xsd:anyURI" use="optional"/>
            <xsd:attribute name="waitSoundMethod" type="http_methods" use="optional"/>
            <xsd:attribute name="digitsMatch" type="xsd:string" use="optional"/>
            <xsd:attribute name="stayAlone" type="booleans" use="optional"/>
            <xsd:attribute name="record" type="booleans" use="optional"/>
            <!-- Deprecated in favor of direction -->
            <xsd:attribute name="recordFileFormat" type="xsd:string" use="optional"/>
            <xsd:attribute name="recordCallbackUrl" type="xsd:anyURI" use="optional"/>
            <xsd:attribute name="recordCallbackMethod" type="http_methods" use="optional"/>
        </xsd:complexType>
    </xsd:element>
    <xsd:element name="Dial">
        <xsd:complexType mixed="true">
            <xsd:choice maxOccurs="unbounded" minOccurs="0">
                <xsd:element ref="Conference"/>
                <xsd:element ref="Number"/>
                <xsd:element ref="Sip"/>
                <xsd:element ref="User"/>
            </xsd:choice>
            <xsd:attribute name="action" type="xsd:anyURI" use="optional"/>
            <xsd:attribute name="method" type="http_methods" use="optional"/>
            <xsd:attribute name="timeout" type="positive_integer" use="optional"/>
            <xsd:attribute name="hangupOnStar" type="booleans" use="optional"/>
            <xsd:attribute name="timeLimit" type="xsd:integer" use="optional"/>
            <xsd:attribute name="callerId" type="xsd:string" use="optional"/>
            <xsd:attribute name="hideCallerId" type="booleans" use="optional"/>
            <xsd:attribute name="callerName" type="xsd:string" use="optional"/>
            <xsd:attribute name="dialMusic" type="xsd:anyURI" use="optional"/>
            <xsd:attribute name="callbackUrl" type="xsd:anyURI" use="optional"/>
            <xsd:attribute name="callbackMethod" type="http_methods" use="optional"/>
            <xsd:attribute name="confirmSound" type="xsd:string" use="optional"/>
            <xsd:attribute name="digitsMatch" type="xsd:string" use="optional"/>
            <xsd:attribute name="straightToVm" type="boolean_or_positive_integer" use="optional"/>
            <xsd:attribute name="heartbeatUrl" type="xsd:anyURI" use="optional"/>
            <xsd:attribute name="heartbeatMethod" type="http_methods" use="optional"/>
            <xsd:attribute name="forwardedFrom" type="xsd:string" use="optional"/>
            <xsd:attribute name="record" type="booleans" use="optional"/>
            <xsd:attribute name="recordDirection" type="zangDirection" use="optional"/>
            <xsd:attribute name="recordCallbackUrl" type="xsd:anyURI" use="optional"/>
            <xsd:attribute name="recordLifetime" type="positive_integer" use="optional"/>
            <xsd:attribute name="recordFormat" type="xsd:string" use="optional"/>
            <xsd:attribute name="ifMachine" type="if_machine" use="optional"/>
            <xsd:attribute name="ifMachineUrl" type="xsd:anyURI" use="optional"/>
            <xsd:attribute name="ifMachineMethod" type="http_methods" use="optional"/>
            <xsd:attribute name="outboundAction" type="booleans" use="optional"/>
        </xsd:complexType>
    </xsd:element>
    <xsd:element name="Gather">
        <xsd:complexType mixed="true">
            <xsd:choice maxOccurs="unbounded" minOccurs="0">
                <xsd:element maxOccurs="unbounded" minOccurs="0" ref="Say"/>
                <xsd:element maxOccurs="unbounded" minOccurs="0" ref="Play"/>
                <xsd:element maxOccurs="unbounded" minOccurs="0" ref="PlayLastRecording"/>
                <xsd:element maxOccurs="unbounded" minOccurs="0" ref="Pause"/>
            </xsd:choice>
            <xsd:attribute name="input" type="gather_input" use="required"/>
            <xsd:attribute name="hints" type="xsd:string" use="optional"/>
            <xsd:attribute name="language" type="bcp_languages" use="optional"/>
            <xsd:attribute name="action" type="xsd:anyURI" use="optional"/>
            <xsd:attribute name="method" type="http_methods" use="optional"/>
            <xsd:attribute name="timeout" type="positive_integer" use="optional"/>
            <xsd:attribute name="finishOnKey" type="xsd:string" use="optional"/>
            <xsd:attribute name="numDigits" type="xsd:integer" use="optional"/>
        </xsd:complexType>
    </xsd:element>
    <xsd:element name="GetSpeech">
        <xsd:complexType mixed="true">
            <xsd:choice maxOccurs="unbounded" minOccurs="0">
                <xsd:element maxOccurs="unbounded" minOccurs="0" ref="Say"/>
                <xsd:element maxOccurs="unbounded" minOccurs="0" ref="Play"/>
                <xsd:element maxOccurs="unbounded" minOccurs="0" ref="PlayLastRecording"/>
                <xsd:element maxOccurs="unbounded" minOccurs="0" ref="Pause"/>
            </xsd:choice>
            <xsd:attribute name="action" type="xsd:anyURI" use="optional"/>
            <xsd:attribute name="method" type="http_methods" use="optional"/>
            <xsd:attribute name="timeout" type="positive_integer" use="optional"/>
            <xsd:attribute name="playBeep" type="booleans" use="optional"/>
            <xsd:attribute name="grammar" type="xsd:anyURI" use="required"/>
        </xsd:complexType>
    </xsd:element>
    <xsd:element name="Hangup">
        <xsd:complexType mixed="true">
            <xsd:attribute name="schedule" type="xsd:integer" use="optional"/>
            <xsd:attribute name="reason" type="xsd:string" use="optional"/>
        </xsd:complexType>
    </xsd:element>
    <xsd:element name="Number">
        <xsd:complexType mixed="true">
            <xsd:attribute name="sendDigits" type="xsd:string" use="optional"/>
            <xsd:attribute name="method" type="http_methods" use="optional"/>
            <xsd:attribute name="url" type="xsd:anyURI" use="optional"/>
            <xsd:attribute name="sendOnPreanswer" type="booleans" use="optional"/>
        </xsd:complexType>
    </xsd:element>
    <xsd:element name="User">
        <xsd:complexType mixed="true">
            <xsd:attribute name="sendDigits" type="xsd:string" use="optional"/>
            <!--      <xsd:attribute name='method' type='http_methods' use='optional' /><xsd:attribute name='url' type='xsd:anyURI' use='optional'/>-->
            <xsd:attribute name="params" type="xsd:string" use="optional"/>
        </xsd:complexType>
    </xsd:element>
    <xsd:element name="Pause">
        <xsd:complexType mixed="true">
            <xsd:attribute name="length" type="xsd:int" use="optional"/>
        </xsd:complexType>
    </xsd:element>
    <xsd:element name="Play" nillable="false">
        <xsd:complexType mixed="true">
            <xsd:attribute name="loop" type="xsd:int" use="optional"/>
            <xsd:attribute name="method" type="http_methods" use="optional"/>
        </xsd:complexType>
    </xsd:element>
    <xsd:element name="Answer"></xsd:element>
    <xsd:element name="PlayLastRecording"/>
    <xsd:element name="Record">
        <xsd:complexType mixed="true">
            <xsd:attribute name="action" type="xsd:anyURI" use="optional"/>
            <xsd:attribute name="method" type="http_methods" use="optional"/>
            <xsd:attribute name="timeout" type="positive_integer" use="optional"/>
            <xsd:attribute name="finishOnKey" type="xsd:string" use="optional"/>
            <xsd:attribute name="maxLength" type="positive_integer" use="optional"/>
            <xsd:attribute name="transcribe" type="booleans" use="optional"/>
            <xsd:attribute name="transcribeCallback" type="xsd:anyURI" use="optional"/>
            <xsd:attribute name="transcribeQuality" type="transcription_quality" use="optional"/>
            <xsd:attribute name="sliceStart" type="positive_integer" use="optional"/>
            <xsd:attribute name="sliceDuration" type="positive_integer" use="optional"/>
            <xsd:attribute name="playBeep" type="booleans" use="optional"/>
            <xsd:attribute name="bothLegs" type="booleans" use="optional"/>
            <!-- Deprecated in favor of direction -->
            <xsd:attribute name="fileFormat" type="xsd:string" use="optional"/>
            <xsd:attribute name="direction" type="zangDirection" use="optional"/>
            <xsd:attribute name="background" type="booleans" use="optional"/>
            <xsd:attribute name="sampleRate" type="sample_rate" use="optional"/>
            <xsd:attribute name="trimSilence" type="booleans" use="optional"/>
            <xsd:attribute name="lifetime" type="positive_integer" use="optional"/>
        </xsd:complexType>
    </xsd:element>
    <xsd:element name="Redirect">
        <xsd:complexType mixed="true">
            <xsd:attribute name="method" type="http_methods" use="optional"/>
        </xsd:complexType>
    </xsd:element>
    <xsd:element name="Ping">
        <xsd:complexType mixed="true">
            <xsd:attribute name="method" type="http_methods" use="optional"/>
        </xsd:complexType>
    </xsd:element>
    <xsd:element name="Reject">
        <xsd:complexType mixed="true">
            <xsd:attribute name="reason" type="reject_modes" use="optional"/>
        </xsd:complexType>
    </xsd:element>
    <xsd:element name="Response">
        <xsd:complexType mixed="true">
            <xsd:choice maxOccurs="unbounded" minOccurs="0">
                <xsd:element maxOccurs="unbounded" minOccurs="0" ref="Say"/>
                <xsd:element maxOccurs="unbounded" minOccurs="0" ref="Play"/>
                <xsd:element maxOccurs="unbounded" minOccurs="0" ref="Answer"/>
                <xsd:element maxOccurs="unbounded" minOccurs="0" ref="PlayLastRecording"/>
                <xsd:element maxOccurs="unbounded" minOccurs="0" ref="Gather"/>
                <xsd:element maxOccurs="unbounded" minOccurs="0" ref="Record"/>
                <xsd:element maxOccurs="unbounded" minOccurs="0" ref="Dial"/>
                <xsd:element maxOccurs="unbounded" minOccurs="0" ref="Hangup"/>
                <xsd:element maxOccurs="unbounded" minOccurs="0" ref="Redirect"/>
                <xsd:element maxOccurs="unbounded" minOccurs="0" ref="Ping"/>
                <xsd:element maxOccurs="unbounded" minOccurs="0" ref="Reject"/>
                <xsd:element maxOccurs="unbounded" minOccurs="0" ref="Pause"/>
                <xsd:element maxOccurs="unbounded" minOccurs="0" ref="Sms"/>
                <xsd:element maxOccurs="unbounded" minOccurs="0" ref="Mms"/>
                <xsd:element maxOccurs="unbounded" minOccurs="0" ref="GetSpeech"/>
            </xsd:choice>
            <xsd:attribute name="statusCallback" type="xsd:anyURI" use="optional"/>
            <xsd:attribute name="statusMethod" type="http_methods" use="optional"/>
            <xsd:attribute name="heartbeatCallback" type="xsd:anyURI" use="optional"/>
            <xsd:attribute name="heartbeatMethod" type="http_methods" use="optional"/>
        </xsd:complexType>
    </xsd:element>
    <xsd:element name="Say">
        <xsd:complexType mixed="true">
            <xsd:attribute name="loop" type="xsd:integer" use="optional"/>
            <xsd:attribute name="voice" type="say_voice" use="optional"/>
            <xsd:attribute name="language" type="tts_languages" use="optional"/>
        </xsd:complexType>
    </xsd:element>
    <xsd:element name="Sms">
        <xsd:complexType mixed="true">
            <xsd:attribute name="to" type="xsd:string" use="required"/>
            <xsd:attribute name="from" type="xsd:string" use="required"/>
            <xsd:attribute name="action" type="xsd:anyURI" use="optional"/>
            <xsd:attribute name="method" type="http_methods" use="optional"/>
            <xsd:attribute name="statusCallback" type="xsd:anyURI" use="optional"/>
            <xsd:attribute name="statusCallbackMethod" type="http_methods" use="optional"/>
        </xsd:complexType>
    </xsd:element>
    <xsd:element name="Mms">
        <xsd:complexType mixed="true">
            <xsd:attribute name="to" type="xsd:string" use="required"/>
            <xsd:attribute name="from" type="xsd:string" use="required"/>
            <xsd:attribute name="action" type="xsd:anyURI" use="optional"/>
            <xsd:attribute name="method" type="http_methods" use="optional"/>
            <xsd:attribute name="statusCallback" type="xsd:anyURI" use="optional"/>
            <xsd:attribute name="mediaUrl" type="xsd:anyURI" use="required"/>
        </xsd:complexType>
    </xsd:element>
    <xsd:element name="Sip">
        <xsd:complexType mixed="true">
            <xsd:attribute name="action" type="xsd:anyURI" use="optional"/>
            <xsd:attribute name="method" type="http_methods" use="optional"/>
            <xsd:attribute name="username" type="xsd:string" use="optional"/>
            <xsd:attribute name="password" type="xsd:string" use="optional"/>
            <xsd:attribute name="sendDigits" type="xsd:string" use="optional"/>
        </xsd:complexType>
    </xsd:element>
</xsd:schema>`

// InboundXML -
type InboundXML struct {
	Response    Response `xml:"Response"`
	RenderCache []byte
}

// GetSchema -
func (i *InboundXML) GetSchema() ([]byte, error) {
	b, err := ioutil.ReadFile("./schema/inboundxml.xsd")
	if err != nil {
		return []byte(ZANG_SCHEMA), nil
	}

	return b, nil
}

// Validate -
func (i *InboundXML) Validate() error {
	rdoc, err := i.Marshal()
	if err != nil {
		return err
	}

	schema, scherr := i.GetSchema()
	if scherr != nil {
		return scherr
	}

	xsdSchema, err := xsd.ParseSchema([]byte(schema))
	if err != nil {
		return err
	}

	doc := golibxml.ParseDoc(rdoc)
	if doc == nil {
		return fmt.Errorf("Could not parse xml document: %v", rdoc)
	}
	defer doc.Free()

	if err := xsdSchema.Validate(xsd.DocPtr(unsafe.Pointer(doc.Ptr))); err != nil {
		return err
	}

	return nil
}

// Marshal -
func (i *InboundXML) Marshal() (string, error) {
	ixml, err := xml.Marshal(i.Response)
	if err != nil {
		return "", err
	}

	// Save cache here so later on we can use the same for .String()
	i.RenderCache = ixml

	return string(ixml), nil
}

// String -
func (i *InboundXML) String() (response string) {
	if i.RenderCache != nil {
		response = string(i.RenderCache)
	}
	return
}

// New -
func New(r Response) (*InboundXML, error) {
	document := &InboundXML{Response: r}

	if err := document.Validate(); err != nil {
		return nil, err
	}

	return document, nil
}
