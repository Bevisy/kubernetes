package kata

import (
	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	_ "k8s.io/apimachinery/pkg/labels"
	"k8s.io/kubernetes/test/e2e/framework"
)

var _ = SIGDescribe("Pod", func() {
	f := framework.NewDefaultFramework("pods")
	var podClient *framework.PodClient
	ginkgo.BeforeEach(func() {
		podClient = f.PodClient()
	})

	ginkgo.It("create a pod using the secure container runtime", func() {
		ginkgo.By("Create Pod")
		runtime := "rune"
		pod := &corev1.Pod{
			ObjectMeta: metav1.ObjectMeta{
				Name: "kata",
				Labels: map[string]string{
					"name": "kata",
				},
			},
			Spec: corev1.PodSpec{
				Containers: []corev1.Container{
					{
						Name:  "nginx",
						Image: "nginx:1.17.3",
					},
				},
				RuntimeClassName: &runtime,
			},
		}
		podClient.Create(pod)

		ginkgo.By("Get the pod")
		podGetting, err := podClient.Get(pod.Name, metav1.GetOptions{})
		framework.ExpectNoError(err, "Failed to get the pod")
		framework.ExpectNoError(f.WaitForPodRunning(podGetting.Name))
		gomega.Expect(podGetting.Name, "kata")

	})
})
