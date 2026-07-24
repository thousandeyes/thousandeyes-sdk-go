package administrative_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/thousandeyes/thousandeyes-sdk-go/v3/administrative"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/client"
)

func TestGetRole(t *testing.T) {
	const (
		roleID = "role-123"
		aid    = "account-group-456"
		token  = "test-token"
	)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %q, want %q", r.Method, http.MethodGet)
		}
		if r.URL.Path != "/roles/"+roleID {
			t.Errorf("path = %q, want %q", r.URL.Path, "/roles/"+roleID)
		}
		if got := r.URL.Query().Get("aid"); got != aid {
			t.Errorf("aid = %q, want %q", got, aid)
		}
		if got := r.Header.Get("Authorization"); got != "Bearer "+token {
			t.Errorf("Authorization = %q, want %q", got, "Bearer "+token)
		}

		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{
			"name": "Organization Admin",
			"roleId": "role-123",
			"isBuiltin": true,
			"permissions": [
				{
					"label": "View users",
					"permissionId": "permission-789",
					"isManagementPermission": true,
					"permission": "users.view"
				}
			]
		}`))
	}))
	t.Cleanup(server.Close)

	config := client.NewConfiguration().
		WithServerUrl(server.URL).
		WithAuthToken(token)
	config.HTTPClient = server.Client()

	apiClient := client.NewAPIClient(config)
	rolesAPI := (*administrative.RolesAPIService)(&apiClient.Common)

	role, response, err := rolesAPI.GetRole(roleID).Aid(aid).Execute()
	if err != nil {
		t.Fatalf("GetRole() error = %v", err)
	}
	if response.StatusCode != http.StatusOK {
		t.Fatalf("GetRole() status = %d, want %d", response.StatusCode, http.StatusOK)
	}
	if got := role.GetRoleId(); got != roleID {
		t.Errorf("GetRole().RoleId = %q, want %q", got, roleID)
	}
	if got := role.GetName(); got != "Organization Admin" {
		t.Errorf("GetRole().Name = %q, want %q", got, "Organization Admin")
	}
	if got := role.GetIsBuiltin(); !got {
		t.Errorf("GetRole().IsBuiltin = %t, want true", got)
	}
	if len(role.Permissions) != 1 {
		t.Fatalf("GetRole().Permissions length = %d, want 1", len(role.Permissions))
	}
	if got := role.Permissions[0].GetPermissionId(); got != "permission-789" {
		t.Errorf("GetRole().Permissions[0].PermissionId = %q, want %q", got, "permission-789")
	}
}

func TestCreateRoleWithDeserializedRequestBody(t *testing.T) {
	const requestJSON = `{
		"name": "Network Operators",
		"permissions": [
			"permission-123",
			"permission-456"
		]
	}`

	var requestBody administrative.RoleRequestBody
	if err := json.Unmarshal([]byte(requestJSON), &requestBody); err != nil {
		t.Fatalf("json.Unmarshal() error = %v", err)
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("method = %q, want %q", r.Method, http.MethodPost)
		}
		if r.URL.Path != "/roles" {
			t.Errorf("path = %q, want %q", r.URL.Path, "/roles")
		}
		if got := r.Header.Get("Content-Type"); got != "application/json" {
			t.Errorf("Content-Type = %q, want %q", got, "application/json")
		}

		var receivedBody administrative.RoleRequestBody
		if err := json.NewDecoder(r.Body).Decode(&receivedBody); err != nil {
			t.Errorf("decoding request body: %v", err)
			http.Error(w, "invalid request body", http.StatusBadRequest)
			return
		}

		if got := receivedBody.GetName(); got != "Network Operators" {
			t.Errorf("request body Name = %q, want %q", got, "Network Operators")
		}
		if got := receivedBody.GetPermissions(); len(got) != 2 ||
			got[0] != "permission-123" || got[1] != "permission-456" {
			t.Errorf(
				"request body Permissions = %v, want %v",
				got,
				[]string{"permission-123", "permission-456"},
			)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write([]byte(`{
			"name": "Network Operators",
			"roleId": "role-123",
			"isBuiltin": false
		}`))
	}))
	t.Cleanup(server.Close)

	config := client.NewConfiguration().WithServerUrl(server.URL)
	config.HTTPClient = server.Client()

	apiClient := client.NewAPIClient(config)
	rolesAPI := (*administrative.RolesAPIService)(&apiClient.Common)

	role, response, err := rolesAPI.CreateRole().RoleRequestBody(requestBody).Execute()
	if err != nil {
		t.Fatalf("CreateRole() error = %v", err)
	}
	if response.StatusCode != http.StatusCreated {
		t.Fatalf("CreateRole() status = %d, want %d", response.StatusCode, http.StatusCreated)
	}
	if got := role.GetRoleId(); got != "role-123" {
		t.Errorf("CreateRole().RoleId = %q, want %q", got, "role-123")
	}
	if got := role.GetName(); got != "Network Operators" {
		t.Errorf("CreateRole().Name = %q, want %q", got, "Network Operators")
	}
}
