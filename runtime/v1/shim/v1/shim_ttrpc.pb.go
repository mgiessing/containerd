// Code generated by protoc-gen-go-ttrpc. DO NOT EDIT.
// source: github.com/containerd/containerd/runtime/v1/shim/v1/shim.proto
package shim

import (
	context "context"
	ttrpc "github.com/containerd/ttrpc"
	empty "github.com/golang/protobuf/ptypes/empty"
)

type ShimService interface {
	State(context.Context, *StateRequest) (*StateResponse, error)
	Create(context.Context, *CreateTaskRequest) (*CreateTaskResponse, error)
	Start(context.Context, *StartRequest) (*StartResponse, error)
	Delete(context.Context, *empty.Empty) (*DeleteResponse, error)
	DeleteProcess(context.Context, *DeleteProcessRequest) (*DeleteResponse, error)
	ListPids(context.Context, *ListPidsRequest) (*ListPidsResponse, error)
	Pause(context.Context, *empty.Empty) (*empty.Empty, error)
	Resume(context.Context, *empty.Empty) (*empty.Empty, error)
	Checkpoint(context.Context, *CheckpointTaskRequest) (*empty.Empty, error)
	Kill(context.Context, *KillRequest) (*empty.Empty, error)
	Exec(context.Context, *ExecProcessRequest) (*empty.Empty, error)
	ResizePty(context.Context, *ResizePtyRequest) (*empty.Empty, error)
	CloseIO(context.Context, *CloseIORequest) (*empty.Empty, error)
	ShimInfo(context.Context, *empty.Empty) (*ShimInfoResponse, error)
	Update(context.Context, *UpdateTaskRequest) (*empty.Empty, error)
	Wait(context.Context, *WaitRequest) (*WaitResponse, error)
}

func RegisterShimService(srv *ttrpc.Server, svc ShimService) {
	srv.RegisterService("containerd.runtime.linux.shim.v1.Shim", &ttrpc.ServiceDesc{
		Methods: map[string]ttrpc.Method{
			"State": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
				var req StateRequest
				if err := unmarshal(&req); err != nil {
					return nil, err
				}
				return svc.State(ctx, &req)
			},
			"Create": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
				var req CreateTaskRequest
				if err := unmarshal(&req); err != nil {
					return nil, err
				}
				return svc.Create(ctx, &req)
			},
			"Start": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
				var req StartRequest
				if err := unmarshal(&req); err != nil {
					return nil, err
				}
				return svc.Start(ctx, &req)
			},
			"Delete": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
				var req empty.Empty
				if err := unmarshal(&req); err != nil {
					return nil, err
				}
				return svc.Delete(ctx, &req)
			},
			"DeleteProcess": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
				var req DeleteProcessRequest
				if err := unmarshal(&req); err != nil {
					return nil, err
				}
				return svc.DeleteProcess(ctx, &req)
			},
			"ListPids": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
				var req ListPidsRequest
				if err := unmarshal(&req); err != nil {
					return nil, err
				}
				return svc.ListPids(ctx, &req)
			},
			"Pause": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
				var req empty.Empty
				if err := unmarshal(&req); err != nil {
					return nil, err
				}
				return svc.Pause(ctx, &req)
			},
			"Resume": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
				var req empty.Empty
				if err := unmarshal(&req); err != nil {
					return nil, err
				}
				return svc.Resume(ctx, &req)
			},
			"Checkpoint": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
				var req CheckpointTaskRequest
				if err := unmarshal(&req); err != nil {
					return nil, err
				}
				return svc.Checkpoint(ctx, &req)
			},
			"Kill": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
				var req KillRequest
				if err := unmarshal(&req); err != nil {
					return nil, err
				}
				return svc.Kill(ctx, &req)
			},
			"Exec": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
				var req ExecProcessRequest
				if err := unmarshal(&req); err != nil {
					return nil, err
				}
				return svc.Exec(ctx, &req)
			},
			"ResizePty": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
				var req ResizePtyRequest
				if err := unmarshal(&req); err != nil {
					return nil, err
				}
				return svc.ResizePty(ctx, &req)
			},
			"CloseIO": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
				var req CloseIORequest
				if err := unmarshal(&req); err != nil {
					return nil, err
				}
				return svc.CloseIO(ctx, &req)
			},
			"ShimInfo": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
				var req empty.Empty
				if err := unmarshal(&req); err != nil {
					return nil, err
				}
				return svc.ShimInfo(ctx, &req)
			},
			"Update": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
				var req UpdateTaskRequest
				if err := unmarshal(&req); err != nil {
					return nil, err
				}
				return svc.Update(ctx, &req)
			},
			"Wait": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
				var req WaitRequest
				if err := unmarshal(&req); err != nil {
					return nil, err
				}
				return svc.Wait(ctx, &req)
			},
		},
	})
}

type shimClient struct {
	client *ttrpc.Client
}

func NewShimClient(client *ttrpc.Client) ShimService {
	return &shimClient{
		client: client,
	}
}

func (c *shimClient) State(ctx context.Context, req *StateRequest) (*StateResponse, error) {
	var resp StateResponse
	if err := c.client.Call(ctx, "containerd.runtime.linux.shim.v1.Shim", "State", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *shimClient) Create(ctx context.Context, req *CreateTaskRequest) (*CreateTaskResponse, error) {
	var resp CreateTaskResponse
	if err := c.client.Call(ctx, "containerd.runtime.linux.shim.v1.Shim", "Create", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *shimClient) Start(ctx context.Context, req *StartRequest) (*StartResponse, error) {
	var resp StartResponse
	if err := c.client.Call(ctx, "containerd.runtime.linux.shim.v1.Shim", "Start", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *shimClient) Delete(ctx context.Context, req *empty.Empty) (*DeleteResponse, error) {
	var resp DeleteResponse
	if err := c.client.Call(ctx, "containerd.runtime.linux.shim.v1.Shim", "Delete", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *shimClient) DeleteProcess(ctx context.Context, req *DeleteProcessRequest) (*DeleteResponse, error) {
	var resp DeleteResponse
	if err := c.client.Call(ctx, "containerd.runtime.linux.shim.v1.Shim", "DeleteProcess", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *shimClient) ListPids(ctx context.Context, req *ListPidsRequest) (*ListPidsResponse, error) {
	var resp ListPidsResponse
	if err := c.client.Call(ctx, "containerd.runtime.linux.shim.v1.Shim", "ListPids", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *shimClient) Pause(ctx context.Context, req *empty.Empty) (*empty.Empty, error) {
	var resp empty.Empty
	if err := c.client.Call(ctx, "containerd.runtime.linux.shim.v1.Shim", "Pause", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *shimClient) Resume(ctx context.Context, req *empty.Empty) (*empty.Empty, error) {
	var resp empty.Empty
	if err := c.client.Call(ctx, "containerd.runtime.linux.shim.v1.Shim", "Resume", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *shimClient) Checkpoint(ctx context.Context, req *CheckpointTaskRequest) (*empty.Empty, error) {
	var resp empty.Empty
	if err := c.client.Call(ctx, "containerd.runtime.linux.shim.v1.Shim", "Checkpoint", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *shimClient) Kill(ctx context.Context, req *KillRequest) (*empty.Empty, error) {
	var resp empty.Empty
	if err := c.client.Call(ctx, "containerd.runtime.linux.shim.v1.Shim", "Kill", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *shimClient) Exec(ctx context.Context, req *ExecProcessRequest) (*empty.Empty, error) {
	var resp empty.Empty
	if err := c.client.Call(ctx, "containerd.runtime.linux.shim.v1.Shim", "Exec", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *shimClient) ResizePty(ctx context.Context, req *ResizePtyRequest) (*empty.Empty, error) {
	var resp empty.Empty
	if err := c.client.Call(ctx, "containerd.runtime.linux.shim.v1.Shim", "ResizePty", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *shimClient) CloseIO(ctx context.Context, req *CloseIORequest) (*empty.Empty, error) {
	var resp empty.Empty
	if err := c.client.Call(ctx, "containerd.runtime.linux.shim.v1.Shim", "CloseIO", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *shimClient) ShimInfo(ctx context.Context, req *empty.Empty) (*ShimInfoResponse, error) {
	var resp ShimInfoResponse
	if err := c.client.Call(ctx, "containerd.runtime.linux.shim.v1.Shim", "ShimInfo", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *shimClient) Update(ctx context.Context, req *UpdateTaskRequest) (*empty.Empty, error) {
	var resp empty.Empty
	if err := c.client.Call(ctx, "containerd.runtime.linux.shim.v1.Shim", "Update", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *shimClient) Wait(ctx context.Context, req *WaitRequest) (*WaitResponse, error) {
	var resp WaitResponse
	if err := c.client.Call(ctx, "containerd.runtime.linux.shim.v1.Shim", "Wait", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
