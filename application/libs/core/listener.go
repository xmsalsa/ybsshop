/**
 * Created by 宅职社 -- mr.wang
 * User: wang
 * Date: 2021-4-15
 * Time: 下午 16:57
 */
package core

type Dispatcher struct {
	listeners map[string][]func(enet *Event)
}

func (this *Dispatcher) AddEventListener(eventName string, callback func(enet *Event)) {
	if _, ok := this.listeners[eventName]; ok {
	} else {
		this.listeners = make(map[string][]func(enet *Event))
	}
	this.listeners[eventName] = append(this.listeners[eventName], callback)
}

func (this *Dispatcher) DispatchEvent(event *Event) {
	data := this.listeners[event.Name]
	//wg.Add(len(data))
	event.Wg.Add(len(data))
	for _, callback := range data {
		go callback(event)
	}
}
