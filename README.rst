zang-go
==========

This golang package is an open source tool built to simplify interaction with
the `Avaya CPaaS <http://www.zang.io>`_ telephony platform. Avaya OneCloud™️ CPaaS  makes adding voice
and SMS to applications fun and easy.

For more information about Avaya CPaaS, please visit:
`Avaya OneCloud™️ CPaaS  <https://www.zang.io/products/cloud>`_

To read the official documentation, please visit: `Avaya CPaaS Docs <http://docs.zang.io/aspx/docs>`_.


Installation
============

Clone the repo, and install via ``go get``:

.. code-block:: bash

    $ go get -u github.com/zang-cloud/zang-go
    
.. code-block:: bash

Usage
======

Authorization
----
In order to authorize against Avaya OneCloud™️ CPaaS  services you'll have to define authentication credentials.

.. code-block:: bash

    export ZANG_CLOUD_ACCOUNT_SID="{YourAccountSid}"
    export ZANG_CLOUD_ACCESS_TOKEN="{YourAccessToken}"
    
.. code-block:: bash

REST
----

See the `Avaya CPaaS REST API documentation <http://docs.zang.io/aspx/rest>`_
for more information.

**NOTE: ** Please go through tests for specific endpoint to see the example

Send SMS Example
----------------

.. code-block:: golang

  package example

  import (
  	"fmt"
  	zang "github.com/zang-cloud/zang-go"
  )

  client, _ := zang.NewClient()

  response, _ := client.SendSms(map[string]string{
    "From": "E.164 Number Format",
    "To":   "E.164 Number Format",
    "Body": "SMS Body",
  })

  fmt.Printf("Send sms response: %+v", response)

.. code-block:: golang

InboundXML
==========

InboundXML is an XML dialect which enables you to control phone call flow.
For more information please visit the `Zang InboundXML documentation
<http://docs.zang.io/aspx/inboundxml>`_.

<Say> Example
-------------

.. code-block:: golang

  ixml, err := New(Response{Say: &Say{
    Voice: "female",
    Value: "Welcome to Avaya CPaaS!",
    Loop:  3,
  }})

  fmt.Print(ixml)

.. code-block:: golang

will render

.. code-block:: xml

    <?xml version="1.0" encoding="UTF-8" standalone="yes"?>
    <Response>
        <Say loop="3" voice="female" language="en">Welcome to Avaya CPaaS!</Say>
    </Response>

.. code-block:: xml
