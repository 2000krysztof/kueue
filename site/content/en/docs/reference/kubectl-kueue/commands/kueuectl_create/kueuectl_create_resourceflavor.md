---
title: kueuectl create resourceflavor
content_type: tool-reference
auto_generated: true
no_list: false
---

<!--
The file is auto-generated from the Go source code of the component using the
[generator](https://github.com/kubernetes-sigs/kueue/tree/main/cmd/kueuectl-docs).
-->

## Synopsis


Create a resource flavor with the given name.

```
kueuectl create resourceflavor NAME [--node-labels KEY=VALUE] [--node-taints KEY[=VALUE]:EFFECT] [--tolerations KEY[=VALUE][:EFFECT]]|:EFFECT [--dry-run STRATEGY]
```


## Examples

```
  # Create a resource flavor
  kueuectl create resourceflavor my-resource-flavor
  
  # Create a resource flavor with labels
  kueuectl create resourceflavor my-resource-flavor \
  --node-labels kubernetes.io/arch=arm64,kubernetes.io/os=linux
  
  # Create a resource flavor with node taints
  kueuectl create resourceflavor my-resource-flavor \
  --node-taints key1=value:NoSchedule,key2:NoExecute
  
  # Create a resource flavor with tolerations
  kueuectl create resourceflavor my-resource-flavor \
  --tolerations key1=value:NoSchedule,key2:NoExecute,key3=value,key4,:PreferNoSchedule
```


## Options


<table style="width: 100%; table-layout: fixed;">
    <colgroup>
        <col span="1" style="width: 10px;" />
        <col span="1" />
    </colgroup>
    <tbody>
    <tr>
        <td colspan="2">--allow-missing-template-keys&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Default: true</td>
    </tr>
    <tr>
        <td></td>
        <td style="line-height: 130%; word-wrap: break-word;">
            <p>If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.</p>
        </td>
    </tr>
    <tr>
        <td colspan="2">-h, --help</td>
    </tr>
    <tr>
        <td></td>
        <td style="line-height: 130%; word-wrap: break-word;">
            <p>help for resourceflavor</p>
        </td>
    </tr>
    <tr>
        <td colspan="2">--node-labels &lt;comma-separated &#39;key=value&#39; pairs&gt;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Default: []</td>
    </tr>
    <tr>
        <td></td>
        <td style="line-height: 130%; word-wrap: break-word;">
            <p>Labels that associate the ResourceFlavor with Nodes that have the same labels.</p>
        </td>
    </tr>
    <tr>
        <td colspan="2">--node-taints strings</td>
    </tr>
    <tr>
        <td></td>
        <td style="line-height: 130%; word-wrap: break-word;">
            <p>Taints that the nodes associated with this ResourceFlavor have.</p>
        </td>
    </tr>
    <tr>
        <td colspan="2">-o, --output string</td>
    </tr>
    <tr>
        <td></td>
        <td style="line-height: 130%; word-wrap: break-word;">
            <p>Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).</p>
        </td>
    </tr>
    <tr>
        <td colspan="2">--show-managed-fields</td>
    </tr>
    <tr>
        <td></td>
        <td style="line-height: 130%; word-wrap: break-word;">
            <p>If true, keep the managedFields when printing objects in JSON or YAML format.</p>
        </td>
    </tr>
    <tr>
        <td colspan="2">--template string</td>
    </tr>
    <tr>
        <td></td>
        <td style="line-height: 130%; word-wrap: break-word;">
            <p>Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview].</p>
        </td>
    </tr>
    <tr>
        <td colspan="2">--tolerations strings</td>
    </tr>
    <tr>
        <td></td>
        <td style="line-height: 130%; word-wrap: break-word;">
            <p>Extra tolerations that will be added to the pods admitted in the quota associated with this resource flavor.</p>
        </td>
    </tr>
    </tbody>
</table>



## Options inherited from parent commands
<table style="width: 100%; table-layout: fixed;">
    <colgroup>
        <col span="1" style="width: 10px;" />
        <col span="1" />
    </colgroup>
    <tbody>
    <tr>
        <td colspan="2">--as string</td>
    </tr>
    <tr>
        <td></td>
        <td style="line-height: 130%; word-wrap: break-word;">
            <p>Username to impersonate for the operation. User could be a regular user or a service account in a namespace.</p>
        </td>
    </tr>
    <tr>
        <td colspan="2">--as-group strings</td>
    </tr>
    <tr>
        <td></td>
        <td style="line-height: 130%; word-wrap: break-word;">
            <p>Group to impersonate for the operation, this flag can be repeated to specify multiple groups.</p>
        </td>
    </tr>
    <tr>
        <td colspan="2">--as-uid string</td>
    </tr>
    <tr>
        <td></td>
        <td style="line-height: 130%; word-wrap: break-word;">
            <p>UID to impersonate for the operation.</p>
        </td>
    </tr>
    <tr>
        <td colspan="2">--cache-dir string&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Default: &#34;$HOME/.kube/cache&#34;</td>
    </tr>
    <tr>
        <td></td>
        <td style="line-height: 130%; word-wrap: break-word;">
            <p>Default cache directory</p>
        </td>
    </tr>
    <tr>
        <td colspan="2">--certificate-authority string</td>
    </tr>
    <tr>
        <td></td>
        <td style="line-height: 130%; word-wrap: break-word;">
            <p>Path to a cert file for the certificate authority</p>
        </td>
    </tr>
    <tr>
        <td colspan="2">--client-certificate string</td>
    </tr>
    <tr>
        <td></td>
        <td style="line-height: 130%; word-wrap: break-word;">
            <p>Path to a client certificate file for TLS</p>
        </td>
    </tr>
    <tr>
        <td colspan="2">--client-key string</td>
    </tr>
    <tr>
        <td></td>
        <td style="line-height: 130%; word-wrap: break-word;">
            <p>Path to a client key file for TLS</p>
        </td>
    </tr>
    <tr>
        <td colspan="2">--cluster string</td>
    </tr>
    <tr>
        <td></td>
        <td style="line-height: 130%; word-wrap: break-word;">
            <p>The name of the kubeconfig cluster to use</p>
        </td>
    </tr>
    <tr>
        <td colspan="2">--context string</td>
    </tr>
    <tr>
        <td></td>
        <td style="line-height: 130%; word-wrap: break-word;">
            <p>The name of the kubeconfig context to use</p>
        </td>
    </tr>
    <tr>
        <td colspan="2">--disable-compression</td>
    </tr>
    <tr>
        <td></td>
        <td style="line-height: 130%; word-wrap: break-word;">
            <p>If true, opt-out of response compression for all requests to the server</p>
        </td>
    </tr>
    <tr>
        <td colspan="2">--dry-run string&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Default: &#34;none&#34;</td>
    </tr>
    <tr>
        <td></td>
        <td style="line-height: 130%; word-wrap: break-word;">
            <p>Must be &#34;none&#34;, &#34;server&#34;, or &#34;client&#34;. If client strategy, only print the object that would be sent, without sending it. If server strategy, submit server-side request without persisting the resource.</p>
        </td>
    </tr>
    <tr>
        <td colspan="2">--insecure-skip-tls-verify</td>
    </tr>
    <tr>
        <td></td>
        <td style="line-height: 130%; word-wrap: break-word;">
            <p>If true, the server&#39;s certificate will not be checked for validity. This will make your HTTPS connections insecure</p>
        </td>
    </tr>
    <tr>
        <td colspan="2">--kubeconfig string</td>
    </tr>
    <tr>
        <td></td>
        <td style="line-height: 130%; word-wrap: break-word;">
            <p>Path to the kubeconfig file to use for CLI requests.</p>
        </td>
    </tr>
    <tr>
        <td colspan="2">-n, --namespace string</td>
    </tr>
    <tr>
        <td></td>
        <td style="line-height: 130%; word-wrap: break-word;">
            <p>If present, the namespace scope for this CLI request</p>
        </td>
    </tr>
    <tr>
        <td colspan="2">--request-timeout string&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Default: &#34;0&#34;</td>
    </tr>
    <tr>
        <td></td>
        <td style="line-height: 130%; word-wrap: break-word;">
            <p>The length of time to wait before giving up on a single server request. Non-zero values should contain a corresponding time unit (e.g. 1s, 2m, 3h). A value of zero means don&#39;t timeout requests.</p>
        </td>
    </tr>
    <tr>
        <td colspan="2">-s, --server string</td>
    </tr>
    <tr>
        <td></td>
        <td style="line-height: 130%; word-wrap: break-word;">
            <p>The address and port of the Kubernetes API server</p>
        </td>
    </tr>
    <tr>
        <td colspan="2">--tls-server-name string</td>
    </tr>
    <tr>
        <td></td>
        <td style="line-height: 130%; word-wrap: break-word;">
            <p>Server name to use for server certificate validation. If it is not provided, the hostname used to contact the server is used</p>
        </td>
    </tr>
    <tr>
        <td colspan="2">--token string</td>
    </tr>
    <tr>
        <td></td>
        <td style="line-height: 130%; word-wrap: break-word;">
            <p>Bearer token for authentication to the API server</p>
        </td>
    </tr>
    <tr>
        <td colspan="2">--user string</td>
    </tr>
    <tr>
        <td></td>
        <td style="line-height: 130%; word-wrap: break-word;">
            <p>The name of the kubeconfig user to use</p>
        </td>
    </tr>
    </tbody>
</table>



## See Also

* [kueuectl create](../)	 - Create a resource

