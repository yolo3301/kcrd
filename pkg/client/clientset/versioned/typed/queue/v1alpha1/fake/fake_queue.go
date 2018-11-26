/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/yolo3301/kcrd/pkg/apis/queue/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeQueues implements QueueInterface
type FakeQueues struct {
	Fake *FakeQueueV1alpha1
	ns   string
}

var queuesResource = schema.GroupVersionResource{Group: "queue.yolo.dev", Version: "v1alpha1", Resource: "queues"}

var queuesKind = schema.GroupVersionKind{Group: "queue.yolo.dev", Version: "v1alpha1", Kind: "Queue"}

// Get takes name of the queue, and returns the corresponding queue object, and an error if there is any.
func (c *FakeQueues) Get(name string, options v1.GetOptions) (result *v1alpha1.Queue, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(queuesResource, c.ns, name), &v1alpha1.Queue{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Queue), err
}

// List takes label and field selectors, and returns the list of Queues that match those selectors.
func (c *FakeQueues) List(opts v1.ListOptions) (result *v1alpha1.QueueList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(queuesResource, queuesKind, c.ns, opts), &v1alpha1.QueueList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.QueueList{ListMeta: obj.(*v1alpha1.QueueList).ListMeta}
	for _, item := range obj.(*v1alpha1.QueueList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested queues.
func (c *FakeQueues) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(queuesResource, c.ns, opts))

}

// Create takes the representation of a queue and creates it.  Returns the server's representation of the queue, and an error, if there is any.
func (c *FakeQueues) Create(queue *v1alpha1.Queue) (result *v1alpha1.Queue, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(queuesResource, c.ns, queue), &v1alpha1.Queue{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Queue), err
}

// Update takes the representation of a queue and updates it. Returns the server's representation of the queue, and an error, if there is any.
func (c *FakeQueues) Update(queue *v1alpha1.Queue) (result *v1alpha1.Queue, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(queuesResource, c.ns, queue), &v1alpha1.Queue{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Queue), err
}

// Delete takes name of the queue and deletes it. Returns an error if one occurs.
func (c *FakeQueues) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(queuesResource, c.ns, name), &v1alpha1.Queue{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeQueues) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(queuesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.QueueList{})
	return err
}

// Patch applies the patch and returns the patched queue.
func (c *FakeQueues) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Queue, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(queuesResource, c.ns, name, data, subresources...), &v1alpha1.Queue{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Queue), err
}
