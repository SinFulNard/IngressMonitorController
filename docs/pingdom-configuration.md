# Pingdom Configuration

## Note
Currently we do not have access to Pingdom account that's why Tests are not verified. Community members having Pingdom account are welcome to contribute in Test Cases.

## Basic
The following properties need to be configured for Pingdom, in addition to the general properties listed 
in the [Configuration section of the README](../README.md#configuration):

| Key      | Description                                      |
|----------|--------------------------------------------------|
| username | Account username for authentication with Pingdom |
| password | Account password for authentication with Pingdom |

## Optional
The following optional property can be included for Pingdom accounts which require multi-user authentication.
More information can be found [Here](https://www.pingdom.com/api/2.1/#multi-user+authentication)

| Key          | Description                                              |
|--------------|----------------------------------------------------------|
| accountEmail | Email account for multi-user authentication with Pingdom |

## Advanced

Currently additional pingdom configurations can be added through a set of annotations to each ingress object, the current supported annotations are:

|                        Annotation                        |                    Description                   |
|:--------------------------------------------------------:|:------------------------------------------------:|
| pingdom.monitor.stakater.com/resolution                  | The pingdom check interval in minutes            |
| pingdom.monitor.stakater.com/send-notification-when-down | How many failed check attempts before notifying  |
| pingdom.monitor.stakater.com/paused                      | Set to "true" to pause checks                    |
| pingdom.monitor.stakater.com/notify-when-back-up         | Set to "false" to disable recovery notifications |
| pingdom.monitor.stakater.com/request-headers             | Custom pingdom request headers (e.g. {"Accept"="application/json"}) |
