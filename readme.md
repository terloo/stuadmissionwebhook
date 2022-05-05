# AdmissionControllerWebhook
编写一个webhook，通过监听两个不同的HTTP路径(validate和mutate)来进行validating和mutating webhook验证

## 请求体，响应体
kube-apiserver会在mutatingWebhook和validatingWebhook两个AdmissionController时回调webhook配置  
请求体为`admission.k8s.io/v1/AdmissionReview`对象，其中Request字段会被赋值  
响应体为同版本的`AdmissionReview`对象，其中Response字段会被赋值  
需要保证Request和Response的uid字段相同

## 清单
```yaml
apiVersion: admissionregistration.k8s.io/v1
kind: ValidatingWebhookConfiguration
metadata:
  name: validating-webhook-sample
  labels:
    apps: stuadmissionwebhook
webhooks:
  # webhook的名字，必须用.分为三段
  - name: stuadmissionwebhook.stuadmissionwebhook.domain
    admissionReviewVersions: ["v1"]
    sideEffects: None
    # 指定该admission作用的资源对象
    rules:
      - apiGroups: ["*"]
        apiVersions: ["*"]
        operations: ["*"]
        resources: ["*"]   # *代表所有资源，*/*代表所有资源及其子资源
    # 指定该admission作用的命名空间的label匹配项
    namespaceSelector:
      matchExpressions:
        - key: app
          operator: In
          values: ["stuadmissionwebhook"]
    clientConfig:
      # 如果服务在集群中，应该配置service字段
      # 如果服务不在集群中，应该配置url字段
      service:
        name: admissionservice
        namespace: stuadmissionwebhook
        # 路由地址
        path: /validate
        # 端口默认为443
        # port: 443
      # https协议所使用的证书
      caBundle: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUMrVENDQWVHZ0F3SUJBZ0lKQU0yN0NvNWdHSUxpTUEwR0NTcUdTSWIzRFFFQkN3VUFNQkl4RURBT0JnTlYKQkFNTUIzZGxZbWh2YjJzd0lCY05Nakl3TlRBME1ESTFNakEwV2hnUE1qRXlNakEwTVRBd01qVXlNRFJhTUJJeApFREFPQmdOVkJBTU1CM2RsWW1odmIyc3dnZ0VpTUEwR0NTcUdTSWIzRFFFQkFRVUFBNElCRHdBd2dnRUtBb0lCCkFRRHVVWXFpSHZVMHJvSnVHcVNTYk4yY3BCR215UXEwZmcxY0cvcGJBdmI1QkZkL2x0VW5aTml4U0Y1OGpUT2MKbjZWV3Flc2JSaldBYUFGbzErRXlEYXBoREM1cVgzZWkwY0I4VFV5N3lvR0FPOTFWWWp5NGpIVThIcitZMFZKWgpMb0lqdWhCTW9na1ZFYlROcVJLdm52L09RTUFQL28rTGt2cytFbDZvdmdHbzhWbUE1Vk41OFc0YkJiaUZMSjZmCnV5bkxRaUhUYWpKa0RVb3dRSkpaN2VyYytLVkdBbmQxOUZiS1NjQ0x0QmpRMWxWc1NEQklkSmltcUxYaUxZRDUKUFVXMWdRRTF3b1NhM2NoUkpMUTZ1Ymxtd0RMeTI2N3NXRnpYMlhiZ2xXTEkrd3BscGZDdCtrQU1LK2t4SzIraApNVG5ubzhCUVJ3ZFdHelhHRmlKTXNzZS9BZ01CQUFHalVEQk9NQjBHQTFVZERnUVdCQlN1K2FlNTVIZmhQTmEyCi9vcG90UHlGNlk1YnpUQWZCZ05WSFNNRUdEQVdnQlN1K2FlNTVIZmhQTmEyL29wb3RQeUY2WTVielRBTUJnTlYKSFJNRUJUQURBUUgvTUEwR0NTcUdTSWIzRFFFQkN3VUFBNElCQVFDUm82RnVjQSt5c1ZaY3hPdTRHQnlKdGpULwpPeWt6L3V2Y0piSTgyaTV3R09uby9HZ1NUYXRCVDZCWXlLeFdoR3ZHV3Z2ZkgwSVB1OU9CWUgvUE1oYVU0MDZ4CjBWanIxc20xQlV3NFFPRmJpczFzZTJjMUQvbGhYcWw3RzhWMGhvRHE1YzhOWUI4cVhTR21pMEx4Zlkvd2Q5TU8KSHN1L3VUeHMwSEFLN1lhcGVDalpzc2FyTk5DdmpwaHdsWTRDdVEzVTBnYmdoeGhVZzlQYy9GT1ZqT1Z4R0REUwpWSHhyTnZ5dmJXaEtLWEl5bk5JdTRLSTVnNGJ2Um9yaEJuZ0xQTDZldDhibit1eDFnekl5aXd5S2QzNEFHQjBFCjRRbEw3emxqOTh1b1hJYS9MelRRbVQ4MGwvMHNxOERoL0s2SXZiYnlkYXc1TktTRU5aL2xWd0ZMYVdqcAotLS0tLUVORCBDRVJUSUZJQ0FURS0tLS0tCg==
```