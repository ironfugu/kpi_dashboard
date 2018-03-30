package kpi_dashboard

import (
	"encoding/json"
	"fmt"

	"github.com/golang/glog"
)

func timeMonthlyHandler(af apiCmd) (*Response, error) {
	result := `{
  "result": [
    {
      "value": [
        {
          "user.device_info.browser.family": "Chrome",
          "result": 30
        },
        {
          "user.device_info.browser.family": "Firefox",
          "result": 15
        },
        {
          "user.device_info.browser.family": "IE",
          "result": 15
        },
        {
          "user.device_info.browser.family": "Mobile Safari",
          "result": 8
        }
      ],
      "timeframe": {
        "start": "2014-05-04T00:00:00.000Z",
        "end": "2014-05-04T01:00:00.000Z"
      }
    },
    {
      "value": [
        {
          "user.device_info.browser.family": "Chrome",
          "result": 20
        },
        {
          "user.device_info.browser.family": "Firefox",
          "result": 19
        },
        {
          "user.device_info.browser.family": "IE",
          "result": 12
        },
        {
          "user.device_info.browser.family": "Mobile Safari",
          "result": 8
        }
      ],
      "timeframe": {
        "start": "2014-05-04T01:00:00.000Z",
        "end": "2014-05-04T02:00:00.000Z"
      }
    },
    {
      "value": [
        {
          "user.device_info.browser.family": "Chrome",
          "result": 28
        },
        {
          "user.device_info.browser.family": "Firefox",
          "result": 12
        },
        {
          "user.device_info.browser.family": "IE",
          "result": 10
        },
        {
          "user.device_info.browser.family": "Mobile Safari",
          "result": 20
        }
      ],
      "timeframe": {
        "start": "2014-05-04T02:00:00.000Z",
        "end": "2014-05-04T03:00:00.000Z"
      }
    },
    {
      "value": [
        {
          "user.device_info.browser.family": "Chrome",
          "result": 25
        },
        {
          "user.device_info.browser.family": "Firefox",
          "result": 13
        },
        {
          "user.device_info.browser.family": "IE",
          "result": 16
        },
        {
          "user.device_info.browser.family": "Mobile Safari",
          "result": 11
        }
      ],
      "timeframe": {
        "start": "2014-05-04T03:00:00.000Z",
        "end": "2014-05-04T04:00:00.000Z"
      }
    },
    {
      "value": [
        {
          "user.device_info.browser.family": "Chrome",
          "result": 17
        },
        {
          "user.device_info.browser.family": "Firefox",
          "result": 14
        },
        {
          "user.device_info.browser.family": "IE",
          "result": 13
        },
        {
          "user.device_info.browser.family": "Mobile Safari",
          "result": 7
        }
      ],
      "timeframe": {
        "start": "2014-05-04T04:00:00.000Z",
        "end": "2014-05-04T05:00:00.000Z"
      }
    },
    {
      "value": [
        {
          "user.device_info.browser.family": "Chrome",
          "result": 11
        },
        {
          "user.device_info.browser.family": "Firefox",
          "result": 18
        },
        {
          "user.device_info.browser.family": "IE",
          "result": 7
        },
        {
          "user.device_info.browser.family": "Mobile Safari",
          "result": 10
        }
      ],
      "timeframe": {
        "start": "2014-05-04T05:00:00.000Z",
        "end": "2014-05-04T06:00:00.000Z"
      }
    },
    {
      "value": [
        {
          "user.device_info.browser.family": "Chrome",
          "result": 17
        },
        {
          "user.device_info.browser.family": "Firefox",
          "result": 12
        },
        {
          "user.device_info.browser.family": "IE",
          "result": 12
        },
        {
          "user.device_info.browser.family": "Mobile Safari",
          "result": 6
        }
      ],
      "timeframe": {
        "start": "2014-05-04T06:00:00.000Z",
        "end": "2014-05-04T07:00:00.000Z"
      }
    },
    {
      "value": [
        {
          "user.device_info.browser.family": "Chrome",
          "result": 0
        },
        {
          "user.device_info.browser.family": "Firefox",
          "result": 0
        },
        {
          "user.device_info.browser.family": "IE",
          "result": 0
        },
        {
          "user.device_info.browser.family": "Mobile Safari",
          "result": 0
        }
      ],
      "timeframe": {
        "start": "2014-05-04T07:00:00.000Z",
        "end": "2014-05-04T08:00:00.000Z"
      }
    },
    {
      "value": [
        {
          "user.device_info.browser.family": "Chrome",
          "result": 16
        },
        {
          "user.device_info.browser.family": "Firefox",
          "result": 15
        },
        {
          "user.device_info.browser.family": "IE",
          "result": 7
        },
        {
          "user.device_info.browser.family": "Mobile Safari",
          "result": 6
        }
      ],
      "timeframe": {
        "start": "2014-05-04T08:00:00.000Z",
        "end": "2014-05-04T09:00:00.000Z"
      }
    },
    {
      "value": [
        {
          "user.device_info.browser.family": "Chrome",
          "result": 13
        },
        {
          "user.device_info.browser.family": "Firefox",
          "result": 7
        },
        {
          "user.device_info.browser.family": "IE",
          "result": 10
        },
        {
          "user.device_info.browser.family": "Mobile Safari",
          "result": 5
        }
      ],
      "timeframe": {
        "start": "2014-05-04T09:00:00.000Z",
        "end": "2014-05-04T10:00:00.000Z"
      }
    },
    {
      "value": [
        {
          "user.device_info.browser.family": "Chrome",
          "result": 18
        },
        {
          "user.device_info.browser.family": "Firefox",
          "result": 7
        },
        {
          "user.device_info.browser.family": "IE",
          "result": 11
        },
        {
          "user.device_info.browser.family": "Mobile Safari",
          "result": 7
        }
      ],
      "timeframe": {
        "start": "2014-05-04T10:00:00.000Z",
        "end": "2014-05-04T11:00:00.000Z"
      }
    },
    {
      "value": [
        {
          "user.device_info.browser.family": "Chrome",
          "result": 14
        },
        {
          "user.device_info.browser.family": "Firefox",
          "result": 15
        },
        {
          "user.device_info.browser.family": "IE",
          "result": 3
        },
        {
          "user.device_info.browser.family": "Mobile Safari",
          "result": 7
        }
      ],
      "timeframe": {
        "start": "2014-05-04T11:00:00.000Z",
        "end": "2014-05-04T12:00:00.000Z"
      }
    },
    {
      "value": [
        {
          "user.device_info.browser.family": "Chrome",
          "result": 30
        },
        {
          "user.device_info.browser.family": "Firefox",
          "result": 18
        },
        {
          "user.device_info.browser.family": "IE",
          "result": 6
        },
        {
          "user.device_info.browser.family": "Mobile Safari",
          "result": 8
        }
      ],
      "timeframe": {
        "start": "2014-05-04T12:00:00.000Z",
        "end": "2014-05-04T13:00:00.000Z"
      }
    },
    {
      "value": [
        {
          "user.device_info.browser.family": "Chrome",
          "result": 21
        },
        {
          "user.device_info.browser.family": "Firefox",
          "result": 12
        },
        {
          "user.device_info.browser.family": "IE",
          "result": 17
        },
        {
          "user.device_info.browser.family": "Mobile Safari",
          "result": 5
        }
      ],
      "timeframe": {
        "start": "2014-05-04T13:00:00.000Z",
        "end": "2014-05-04T14:00:00.000Z"
      }
    },
    {
      "value": [
        {
          "user.device_info.browser.family": "Chrome",
          "result": 25
        },
        {
          "user.device_info.browser.family": "Firefox",
          "result": 22
        },
        {
          "user.device_info.browser.family": "IE",
          "result": 11
        },
        {
          "user.device_info.browser.family": "Mobile Safari",
          "result": 11
        }
      ],
      "timeframe": {
        "start": "2014-05-04T14:00:00.000Z",
        "end": "2014-05-04T15:00:00.000Z"
      }
    },
    {
      "value": [
        {
          "user.device_info.browser.family": "Chrome",
          "result": 34
        },
        {
          "user.device_info.browser.family": "Firefox",
          "result": 24
        },
        {
          "user.device_info.browser.family": "IE",
          "result": 12
        },
        {
          "user.device_info.browser.family": "Mobile Safari",
          "result": 10
        }
      ],
      "timeframe": {
        "start": "2014-05-04T15:00:00.000Z",
        "end": "2014-05-04T16:00:00.000Z"
      }
    },
    {
      "value": [
        {
          "user.device_info.browser.family": "Chrome",
          "result": 24
        },
        {
          "user.device_info.browser.family": "Firefox",
          "result": 30
        },
        {
          "user.device_info.browser.family": "IE",
          "result": 19
        },
        {
          "user.device_info.browser.family": "Mobile Safari",
          "result": 9
        }
      ],
      "timeframe": {
        "start": "2014-05-04T16:00:00.000Z",
        "end": "2014-05-04T17:00:00.000Z"
      }
    },
    {
      "value": [
        {
          "user.device_info.browser.family": "Chrome",
          "result": 31
        },
        {
          "user.device_info.browser.family": "Firefox",
          "result": 25
        },
        {
          "user.device_info.browser.family": "IE",
          "result": 13
        },
        {
          "user.device_info.browser.family": "Mobile Safari",
          "result": 14
        }
      ],
      "timeframe": {
        "start": "2014-05-04T17:00:00.000Z",
        "end": "2014-05-04T18:00:00.000Z"
      }
    },
    {
      "value": [
        {
          "user.device_info.browser.family": "Chrome",
          "result": 36
        },
        {
          "user.device_info.browser.family": "Firefox",
          "result": 22
        },
        {
          "user.device_info.browser.family": "IE",
          "result": 18
        },
        {
          "user.device_info.browser.family": "Mobile Safari",
          "result": 12
        }
      ],
      "timeframe": {
        "start": "2014-05-04T18:00:00.000Z",
        "end": "2014-05-04T19:00:00.000Z"
      }
    },
    {
      "value": [
        {
          "user.device_info.browser.family": "Chrome",
          "result": 32
        },
        {
          "user.device_info.browser.family": "Firefox",
          "result": 37
        },
        {
          "user.device_info.browser.family": "IE",
          "result": 17
        },
        {
          "user.device_info.browser.family": "Mobile Safari",
          "result": 7
        }
      ],
      "timeframe": {
        "start": "2014-05-04T19:00:00.000Z",
        "end": "2014-05-04T20:00:00.000Z"
      }
    },
    {
      "value": [
        {
          "user.device_info.browser.family": "Chrome",
          "result": 33
        },
        {
          "user.device_info.browser.family": "Firefox",
          "result": 28
        },
        {
          "user.device_info.browser.family": "IE",
          "result": 12
        },
        {
          "user.device_info.browser.family": "Mobile Safari",
          "result": 10
        }
      ],
      "timeframe": {
        "start": "2014-05-04T20:00:00.000Z",
        "end": "2014-05-04T21:00:00.000Z"
      }
    },
    {
      "value": [
        {
          "user.device_info.browser.family": "Chrome",
          "result": 37
        },
        {
          "user.device_info.browser.family": "Firefox",
          "result": 28
        },
        {
          "user.device_info.browser.family": "IE",
          "result": 16
        },
        {
          "user.device_info.browser.family": "Mobile Safari",
          "result": 13
        }
      ],
      "timeframe": {
        "start": "2014-05-04T21:00:00.000Z",
        "end": "2014-05-04T22:00:00.000Z"
      }
    },
    {
      "value": [
        {
          "user.device_info.browser.family": "Chrome",
          "result": 35
        },
        {
          "user.device_info.browser.family": "Firefox",
          "result": 25
        },
        {
          "user.device_info.browser.family": "IE",
          "result": 14
        },
        {
          "user.device_info.browser.family": "Mobile Safari",
          "result": 11
        }
      ],
      "timeframe": {
        "start": "2014-05-04T22:00:00.000Z",
        "end": "2014-05-04T23:00:00.000Z"
      }
    },
    {
      "value": [
        {
          "user.device_info.browser.family": "Chrome",
          "result": 27
        },
        {
          "user.device_info.browser.family": "Firefox",
          "result": 30
        },
        {
          "user.device_info.browser.family": "IE",
          "result": 18
        },
        {
          "user.device_info.browser.family": "Mobile Safari",
          "result": 8
        }
      ],
      "timeframe": {
        "start": "2014-05-04T23:00:00.000Z",
        "end": "2014-05-05T00:00:00.000Z"
      }
    }
  ]
}`
	var r interface{}
	err := json.Unmarshal([]byte(result), &r)
	if err != nil {
		err := fmt.Errorf("could not marshal result: %+v", err)
		glog.Error(err)
		return &Response{Error: &ErrorResp{Reason: err.Error()}}, nil
	}
	return &Response{Result: r}, nil
}
