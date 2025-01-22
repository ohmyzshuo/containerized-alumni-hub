package alumni_test

import (
    "alumni_hub/internal/alumni"
    _ "alumni_hub/internal/utils"
    "bytes"
    "encoding/json"
    "errors"
    "fmt"
    "github.com/gin-gonic/gin"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/mock"
    "net/http"
    "net/http/httptest"
    "strconv"
    "strings"
    "testing"
)

type MockService struct {
    mock.Mock
}

func (m *MockService) GetAlumnus(id int) (*alumni.Alumni, error) {
    args := m.Called(id)
    if args.Get(0) == nil {
        return nil, args.Error(1)
    }
    return args.Get(0).(*alumni.Alumni), args.Error(1)
}

func (m *MockService) GetAlumni(page, pageSize int, searchQuery string) ([]alumni.Alumni, int64, error) {
    args := m.Called(page, pageSize, searchQuery)
    return args.Get(0).([]alumni.Alumni), args.Get(1).(int64), args.Error(2)
}

func (m *MockService) CreateAlumni(alu *alumni.Alumni) (*alumni.Alumni, error) {
    args := m.Called(alu)
    return args.Get(0).(*alumni.Alumni), args.Error(1)
}

func (m *MockService) DeleteAlumni(id int) error {
    args := m.Called(id)
    return args.Error(0)
}

func (m *MockService) CheckAlumniExistence(matricNo string) (alumni.Alumni, bool, error) {
    args := m.Called(matricNo)
    return args.Get(0).(alumni.Alumni), args.Bool(1), args.Error(2)
}

func (m *MockService) GetAlumniByEmail(email string) (*alumni.Alumni, error) {
    args := m.Called(email)
    if args.Get(0) != nil {
        return args.Get(0).(*alumni.Alumni), args.Error(1)
    }
    return nil, args.Error(1)
}

func (m *MockService) GetAlumniByMatricNo(matricNo string) (*alumni.Alumni, error) {
    args := m.Called(matricNo)
    if args.Get(0) != nil {
        return args.Get(0).(*alumni.Alumni), args.Error(1)
    }
    return nil, args.Error(1)
}
func (m *MockService) GetAlumniByToken(token string) (*alumni.Alumni, error) {
    args := m.Called(token)
    if args.Get(0) != nil {
        return args.Get(0).(*alumni.Alumni), args.Error(1)
    }
    return nil, args.Error(1)
}
func (m *MockService) GetAlumnusByID(id uint) (*alumni.Alumni, error) {
    args := m.Called(id)
    if args.Get(0) == nil {
        return nil, args.Error(1)
    }
    return args.Get(0).(*alumni.Alumni), args.Error(1)
}

func (m *MockService) UpdateAlumni(id uint, alumnus *alumni.Alumni) (*alumni.Alumni, error) {
    args := m.Called(id, alumnus)
    if args.Get(0) == nil {
        return nil, args.Error(1)
    }
    return args.Get(0).(*alumni.Alumni), args.Error(1)
}

func TestGetAlumniHandler(t *testing.T) {
    gin.SetMode(gin.TestMode)

    tests := []struct {
        name         string
        page         int
        pageSize     int
        searchQuery  string
        mockAlumni   []alumni.Alumni
        mockTotal    int64
        expectedCode int
        expectedMeta gin.H
    }{
        {
            name:        "Success with default pagination",
            page:        1,
            pageSize:    15,
            searchQuery: "",
            mockAlumni: []alumni.Alumni{
                {ID: 1, Name: "John Doe", Password: "secret"},
                {ID: 2, Name: "Jane Doe", Password: "secret"},
            },
            mockTotal:    2,
            expectedCode: http.StatusOK,
            expectedMeta: gin.H{
                "page":        1,
                "page_size":   15,
                "total":       int64(2),
                "total_pages": int64(1),
            },
        },
        {
            name:        "Success with custom pagination",
            page:        2,
            pageSize:    10,
            searchQuery: "John",
            mockAlumni: []alumni.Alumni{
                {ID: 1, Name: "John Doe", Password: "secret"},
            },
            mockTotal:    11,
            expectedCode: http.StatusOK,
            expectedMeta: gin.H{
                "page":        2,
                "page_size":   10,
                "total":       int64(11),
                "total_pages": int64(2),
            },
        },
        {
            name:        "Invalid page number defaults to 1",
            page:        -1,
            pageSize:    15,
            searchQuery: "",
            mockAlumni: []alumni.Alumni{
                {ID: 1, Name: "John Doe", Password: "secret"},
            },
            mockTotal:    1,
            expectedCode: http.StatusOK,
            expectedMeta: gin.H{
                "page":        1,
                "page_size":   15,
                "total":       int64(1),
                "total_pages": int64(1),
            },
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            mockService := new(MockService)
            handler := alumni.NewHandler(mockService)

            expectedPage := tt.page
            if expectedPage < 1 {
                expectedPage = 1
            }
            expectedPageSize := tt.pageSize
            if expectedPageSize < 1 {
                expectedPageSize = 15
            }
            mockService.On("GetAlumni", expectedPage, expectedPageSize, tt.searchQuery).
                Return(tt.mockAlumni, tt.mockTotal, nil)

            w := httptest.NewRecorder()
            c, _ := gin.CreateTestContext(w)

            url := fmt.Sprintf("/alumni?page=%d&pageSize=%d", tt.page, tt.pageSize)
            if tt.searchQuery != "" {
                url += "&search=" + tt.searchQuery
            }
            c.Request, _ = http.NewRequest("GET", url, nil)
            c.Request.URL.RawQuery = url[strings.Index(url, "?")+1:]

            handler.GetAlumni(c)

            assert.Equal(t, tt.expectedCode, w.Code)

            var response struct {
                Code    int             `json:"code"`
                Message string          `json:"message"`
                Data    []alumni.Alumni `json:"data"`
                Meta    gin.H           `json:"meta"`
            }
            err := json.Unmarshal(w.Body.Bytes(), &response)
            assert.NoError(t, err)

            assert.Equal(t, tt.expectedCode, response.Code)
            assert.Equal(t, "Success", response.Message)

            for _, a := range response.Data {
                assert.Empty(t, a.Password, "Password should be empty in response")
            }

            assert.Equal(t, float64(tt.expectedMeta["page"].(int)), response.Meta["page"].(float64))
            assert.Equal(t, float64(tt.expectedMeta["page_size"].(int)), response.Meta["page_size"].(float64))
            assert.Equal(t, float64(tt.expectedMeta["total"].(int64)), response.Meta["total"].(float64))
            assert.Equal(t, float64(tt.expectedMeta["total_pages"].(int64)), response.Meta["total_pages"].(float64))

            mockService.AssertExpectations(t)
        })
    }
}

func TestUpdateAlumniHandler(t *testing.T) {
    gin.SetMode(gin.TestMode)

    mockService := new(MockService)
    handler := alumni.NewHandler(mockService)

    w := httptest.NewRecorder()
    c, _ := gin.CreateTestContext(w)

    alumnusID := uint(1)
    updatedAlumnus := &alumni.Alumni{
        Name:     "Updated Name",
        Email:    "updated@example.com",
        MatricNo: "12345",
    }
    returnedAlumnus := &alumni.Alumni{
        ID:       alumnusID,
        Name:     updatedAlumnus.Name,
        Email:    updatedAlumnus.Email,
        MatricNo: updatedAlumnus.MatricNo,
    }

    mockService.On("UpdateAlumni", alumnusID, updatedAlumnus).Return(returnedAlumnus, nil)

    jsonValue, _ := json.Marshal(updatedAlumnus)
    c.Request, _ = http.NewRequest("PUT", "/alumni/1", bytes.NewBuffer(jsonValue))
    c.Request.Header.Set("Content-Type", "application/json")
    c.Params = gin.Params{gin.Param{Key: "id", Value: "1"}}

    handler.UpdateAlumni(c)

    assert.Equal(t, http.StatusOK, w.Code)
    var response map[string]interface{}
    json.Unmarshal(w.Body.Bytes(), &response)
    assert.Equal(t, float64(http.StatusOK), response["code"])
    assert.Equal(t, "Alumni updated successfully", response["message"])
    assert.NotNil(t, response["data"])

    mockService.AssertExpectations(t)
}

func TestDeleteAlumniHandler(t *testing.T) {
    gin.SetMode(gin.TestMode)

    mockService := new(MockService)
    handler := alumni.NewHandler(mockService)

    w := httptest.NewRecorder()
    c, _ := gin.CreateTestContext(w)

    alumnusID := 1
    mockService.On("DeleteAlumni", alumnusID).Return(nil)

    c.Params = gin.Params{gin.Param{Key: "id", Value: strconv.Itoa(alumnusID)}}
    handler.DeleteAlumni(c)

    assert.Equal(t, http.StatusOK, w.Code)
    var response map[string]interface{}
    json.Unmarshal(w.Body.Bytes(), &response)
    assert.Equal(t, float64(http.StatusOK), response["code"])
    assert.Equal(t, "Alumni hidden successfully", response["message"])

    mockService.AssertExpectations(t)
}

func TestCheckAlumniExistenceHandler(t *testing.T) {
    gin.SetMode(gin.TestMode)

    tests := []struct {
        name           string
        matricNo       string
        mockReturn     alumni.Alumni
        mockExists     bool
        mockError      error
        expectedStatus int
        expectedBody   map[string]interface{}
    }{
        {
            name:     "Alumni exists",
            matricNo: "12345",
            mockReturn: alumni.Alumni{
                Name:     "John Doe",
                MatricNo: "12345",
            },
            mockExists:     true,
            mockError:      nil,
            expectedStatus: http.StatusOK,
            expectedBody: map[string]interface{}{
                "code": http.StatusOK,
                "data": map[string]interface{}{
                    "name":      "John Doe",
                    "matric_no": "12345",
                },
                "message": "The matric number exists in the database, please log in with the same username",
            },
        },
        {
            name:           "Alumni does not exist",
            matricNo:       "67890",
            mockReturn:     alumni.Alumni{},
            mockExists:     false,
            mockError:      nil,
            expectedStatus: http.StatusOK,
            expectedBody: map[string]interface{}{
                "code":    http.StatusOK,
                "data":    nil,
                "message": "The matric number doesn’t exist in the database, please contact staff",
            },
        },
        {
            name:           "Service error",
            matricNo:       "error",
            mockReturn:     alumni.Alumni{},
            mockExists:     false,
            mockError:      assert.AnError,
            expectedStatus: http.StatusInternalServerError,
            expectedBody: map[string]interface{}{
                "code":    http.StatusInternalServerError,
                "data":    nil,
                "message": "Error checking alumni existence",
            },
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            mockSvc := new(MockService)
            mockSvc.On("CheckAlumniExistence", tt.matricNo).Return(tt.mockReturn, tt.mockExists, tt.mockError)

            // Create handler using the NewHandler function
            handler := alumni.NewHandler(mockSvc)

            // Create a request to pass to our handler
            req, err := http.NewRequest(http.MethodGet, "/check-alumni-existence", nil)
            assert.NoError(t, err)

            // Add query parameter
            q := req.URL.Query()
            q.Add("matric_no", tt.matricNo)
            req.URL.RawQuery = q.Encode()

            // Create a response recorder to record the response
            w := httptest.NewRecorder()

            // Create a new gin context
            c, _ := gin.CreateTestContext(w)
            c.Request = req

            // Call the handler function
            handler.CheckAlumniExistence(c)

            // Assert the response status code
            assert.Equal(t, tt.expectedStatus, w.Code)

            // Assert the response body
            var responseBody map[string]interface{}
            err = json.Unmarshal(w.Body.Bytes(), &responseBody)
            assert.NoError(t, err)

            // Convert code to int for comparison
            if code, ok := responseBody["code"].(float64); ok {
                responseBody["code"] = int(code)
            }

            // Extract the relevant fields from the actual data for comparison
            if data, ok := responseBody["data"].(map[string]interface{}); ok {
                if tt.mockExists {
                    responseBody["data"] = map[string]interface{}{
                        "name":      data["name"],
                        "matric_no": data["matric_no"],
                    }
                } else {
                    responseBody["data"] = nil
                }
            }

            assert.Equal(t, tt.expectedBody, responseBody)
        })
    }
}

func TestGetAlumniByEmailHandler(t *testing.T) {
    gin.SetMode(gin.TestMode)

    tests := []struct {
        name           string
        email          string
        mockReturn     *alumni.Alumni
        mockError      error
        expectedStatus int
        expectedBody   map[string]interface{}
    }{
        {
            name:  "Alumni found",
            email: "john.doe@example.com",
            mockReturn: &alumni.Alumni{
                Name:     "John Doe",
                Email:    "john.doe@example.com",
                MatricNo: "12345",
            },
            mockError:      nil,
            expectedStatus: http.StatusOK,
            expectedBody: map[string]interface{}{
                "message": "success",
                "data": map[string]interface{}{
                    "name":      "John Doe",
                    "email":     "john.doe@example.com",
                    "matric_no": "12345",
                },
                "code": http.StatusOK,
            },
        },
        {
            name:           "Alumni not found",
            email:          "notfound@example.com",
            mockReturn:     nil,
            mockError:      nil,
            expectedStatus: http.StatusNotFound,
            expectedBody: map[string]interface{}{
                "error": "Alumni not found",
            },
        },
        {
            name:           "Service error",
            email:          "error@example.com",
            mockReturn:     nil,
            mockError:      assert.AnError,
            expectedStatus: http.StatusInternalServerError,
            expectedBody: map[string]interface{}{
                "error": assert.AnError.Error(),
            },
        },
        {
            name:           "Invalid request body",
            email:          "",
            mockReturn:     nil,
            mockError:      nil,
            expectedStatus: http.StatusBadRequest,
            expectedBody: map[string]interface{}{
                "error": "Invalid request body",
            },
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            mockSvc := new(MockService)
            mockSvc.On("GetAlumniByEmail", tt.email).Return(tt.mockReturn, tt.mockError)

            // Create handler using the NewHandler function
            handler := alumni.NewHandler(mockSvc)

            // Create a request to pass to our handler
            var req *http.Request
            var err error

            if tt.name == "Invalid request body" {
                req, err = http.NewRequest(http.MethodPost, "/get-alumni-by-email", nil)
            } else {
                reqBody := map[string]string{"email": tt.email}
                reqBodyBytes, _ := json.Marshal(reqBody)
                req, err = http.NewRequest(http.MethodPost, "/get-alumni-by-email", bytes.NewBuffer(reqBodyBytes))
            }

            assert.NoError(t, err)

            // Create a response recorder to record the response
            w := httptest.NewRecorder()

            // Create a new gin context
            c, _ := gin.CreateTestContext(w)
            c.Request = req

            // Call the handler function
            handler.GetAlumniByEmail(c)

            // Assert the response status code
            assert.Equal(t, tt.expectedStatus, w.Code)

            // Assert the response body
            var responseBody map[string]interface{}
            err = json.Unmarshal(w.Body.Bytes(), &responseBody)
            assert.NoError(t, err)

            // Convert code to int for comparison if it exists
            if code, ok := responseBody["code"].(float64); ok {
                responseBody["code"] = int(code)
            }

            // Remove password field for comparison if it exists
            if data, ok := responseBody["data"].(map[string]interface{}); ok {
                delete(data, "password")
            }

            // Remove unexpected fields from response body for comparison
            if data, ok := responseBody["data"].(map[string]interface{}); ok {
                for key := range data {
                    if _, expected := tt.expectedBody["data"].(map[string]interface{})[key]; !expected {
                        delete(data, key)
                    }
                }
            }

            assert.Equal(t, tt.expectedBody, responseBody)
        })
    }
}

func TestGetAlumniByMatricNoHandler(t *testing.T) {
    gin.SetMode(gin.TestMode)

    tests := []struct {
        name           string
        matricNo       string
        mockReturn     *alumni.Alumni
        mockError      error
        expectedStatus int
        expectedBody   map[string]interface{}
    }{
        {
            name:     "Alumni found",
            matricNo: "12345",
            mockReturn: &alumni.Alumni{
                Name:     "John Doe",
                Email:    "john.doe@example.com",
                MatricNo: "12345",
            },
            mockError:      nil,
            expectedStatus: http.StatusOK,
            expectedBody: map[string]interface{}{
                "message": "success",
                "data": map[string]interface{}{
                    "name":      "John Doe",
                    "email":     "john.doe@example.com",
                    "matric_no": "12345",
                },
                "code": http.StatusOK,
            },
        },
        {
            name:           "Alumni not found",
            matricNo:       "67890",
            mockReturn:     nil,
            mockError:      nil,
            expectedStatus: http.StatusNotFound,
            expectedBody: map[string]interface{}{
                "error": "Alumni not found",
            },
        },
        {
            name:           "Service error",
            matricNo:       "error",
            mockReturn:     nil,
            mockError:      assert.AnError,
            expectedStatus: http.StatusInternalServerError,
            expectedBody: map[string]interface{}{
                "error": assert.AnError.Error(),
            },
        },
        {
            name:           "Invalid request body",
            matricNo:       "",
            mockReturn:     nil,
            mockError:      nil,
            expectedStatus: http.StatusBadRequest,
            expectedBody: map[string]interface{}{
                "error": "Invalid request body",
            },
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            mockSvc := new(MockService)
            mockSvc.On("GetAlumniByMatricNo", tt.matricNo).Return(tt.mockReturn, tt.mockError)

            // Create handler using the NewHandler function
            handler := alumni.NewHandler(mockSvc)

            // Create a request to pass to our handler
            var req *http.Request
            var err error

            if tt.name == "Invalid request body" {
                req, err = http.NewRequest(http.MethodPost, "/get-alumni-by-matric-no", nil)
            } else {
                reqBody := map[string]string{"matric_no": tt.matricNo}
                reqBodyBytes, _ := json.Marshal(reqBody)
                req, err = http.NewRequest(http.MethodPost, "/get-alumni-by-matric-no", bytes.NewBuffer(reqBodyBytes))
            }

            assert.NoError(t, err)

            // Create a response recorder to record the response
            w := httptest.NewRecorder()

            // Create a new gin context
            c, _ := gin.CreateTestContext(w)
            c.Request = req

            // Call the handler function
            handler.GetAlumniByMatricNo(c)

            // Assert the response status code
            assert.Equal(t, tt.expectedStatus, w.Code)

            // Assert the response body
            var responseBody map[string]interface{}
            err = json.Unmarshal(w.Body.Bytes(), &responseBody)
            assert.NoError(t, err)

            // Convert code to int for comparison if it exists
            if code, ok := responseBody["code"].(float64); ok {
                responseBody["code"] = int(code)
            }

            // Remove password field for comparison if it exists
            if data, ok := responseBody["data"].(map[string]interface{}); ok {
                delete(data, "password")
            }

            // Remove unexpected fields from response body for comparison
            if data, ok := responseBody["data"].(map[string]interface{}); ok {
                for key := range data {
                    if _, expected := tt.expectedBody["data"].(map[string]interface{})[key]; !expected {
                        delete(data, key)
                    }
                }
            }

            assert.Equal(t, tt.expectedBody, responseBody)
        })
    }
}

func TestGetAlumniByTokenHandler(t *testing.T) {
    gin.SetMode(gin.TestMode)

    tests := []struct {
        name           string
        token          string
        mockReturn     *alumni.Alumni
        mockError      error
        expectedStatus int
        expectedBody   map[string]interface{}
    }{
        {
            name:  "User found",
            token: "valid-token",
            mockReturn: &alumni.Alumni{
                Name:     "John Doe",
                Email:    "john.doe@example.com",
                MatricNo: "12345",
            },
            mockError:      nil,
            expectedStatus: http.StatusOK,
            expectedBody: map[string]interface{}{
                "data": map[string]interface{}{
                    "name":      "John Doe",
                    "email":     "john.doe@example.com",
                    "matric_no": "12345",
                },
                "code": http.StatusOK,
            },
        },
        {
            name:           "User not found",
            token:          "invalid-token",
            mockReturn:     nil,
            mockError:      nil,
            expectedStatus: http.StatusNotFound,
            expectedBody: map[string]interface{}{
                "message": "User not found",
            },
        },
        {
            name:           "Service error",
            token:          "error-token",
            mockReturn:     nil,
            mockError:      assert.AnError,
            expectedStatus: http.StatusInternalServerError,
            expectedBody: map[string]interface{}{
                "error": assert.AnError.Error(),
            },
        },
        {
            name:           "Missing token",
            token:          "",
            mockReturn:     nil,
            mockError:      nil,
            expectedStatus: http.StatusBadRequest,
            expectedBody: map[string]interface{}{
                "code":    http.StatusBadRequest,
                "message": "UserToken is required",
                "data":    nil,
            },
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            mockSvc := new(MockService)
            mockSvc.On("GetAlumniByToken", tt.token).Return(tt.mockReturn, tt.mockError)

            // Create handler using the NewHandler function
            handler := alumni.NewHandler(mockSvc)

            // Create a request to pass to our handler
            req, err := http.NewRequest(http.MethodGet, "/get-alumni-by-token", nil)
            assert.NoError(t, err)

            if tt.token != "" {
                req.Header.Set("Authorization", "Bearer "+tt.token)
            }

            // Create a response recorder to record the response
            w := httptest.NewRecorder()

            // Create a new gin context
            c, _ := gin.CreateTestContext(w)
            c.Request = req

            // Call the handler function
            handler.GetAlumniByToken(c)

            // Assert the response status code
            assert.Equal(t, tt.expectedStatus, w.Code)

            // Assert the response body
            var responseBody map[string]interface{}
            err = json.Unmarshal(w.Body.Bytes(), &responseBody)
            assert.NoError(t, err)

            // Convert code to int for comparison if it exists
            if code, ok := responseBody["code"].(float64); ok {
                responseBody["code"] = int(code)
            }

            // Remove password field for comparison if it exists
            if data, ok := responseBody["data"].(map[string]interface{}); ok {
                delete(data, "password")
            }

            // Remove unexpected fields from response body for comparison
            if data, ok := responseBody["data"].(map[string]interface{}); ok {
                for key := range data {
                    if _, expected := tt.expectedBody["data"].(map[string]interface{})[key]; !expected {
                        delete(data, key)
                    }
                }
            }

            assert.Equal(t, tt.expectedBody, responseBody)
        })
    }
}

func TestResetAlumnusPasswordHandler(t *testing.T) {
    // Setup
    gin.SetMode(gin.TestMode)

    tests := []struct {
        name           string
        id             string
        setupMock      func(*MockService)
        expectedStatus int
        expectedMsg    string
    }{
        {
            name: "Success",
            id:   "1",
            setupMock: func(ms *MockService) {
                alumnus := &alumni.Alumni{
                    MatricNo: "12345",
                }
                ms.On("GetAlumnusByID", uint(1)).Return(alumnus, nil)
                ms.On("UpdateAlumni", uint(1), mock.MatchedBy(func(a *alumni.Alumni) bool {
                    return true // You can add more specific matching criteria if needed
                })).Return(alumnus, nil)
            },
            expectedStatus: http.StatusOK,
            expectedMsg:    "Password reset successfully",
        },
        {
            name: "Invalid ID",
            id:   "invalid",
            setupMock: func(ms *MockService) {
                // No mock setup needed for invalid ID
            },
            expectedStatus: http.StatusBadRequest,
            expectedMsg:    "Invalid alumnus ID",
        },
        {
            name: "Get Alumnus Error",
            id:   "1",
            setupMock: func(ms *MockService) {
                ms.On("GetAlumnusByID", uint(1)).Return(nil, errors.New("database error"))
            },
            expectedStatus: http.StatusInternalServerError,
            expectedMsg:    "database error",
        },
        {
            name: "Update Error",
            id:   "1",
            setupMock: func(ms *MockService) {
                alumnus := &alumni.Alumni{
                    MatricNo: "12345",
                }
                ms.On("GetAlumnusByID", uint(1)).Return(alumnus, nil)
                ms.On("UpdateAlumni", uint(1), mock.MatchedBy(func(a *alumni.Alumni) bool {
                    return true
                })).Return(nil, errors.New("update error"))
            },
            expectedStatus: http.StatusInternalServerError,
            expectedMsg:    "update error",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            // Create mock service and setup expectations
            mockService := new(MockService)
            if tt.setupMock != nil {
                tt.setupMock(mockService)
            }

            // Create handler using NewHandler constructor
            handler := alumni.NewHandler(mockService)

            // Create test context
            w := httptest.NewRecorder()
            c, _ := gin.CreateTestContext(w)
            c.Params = gin.Params{
                {Key: "id", Value: tt.id},
            }

            // Call the handler
            handler.ResetAlumnusPassword(c)

            // Assert response
            assert.Equal(t, tt.expectedStatus, w.Code)

            var response map[string]interface{}
            err := json.Unmarshal(w.Body.Bytes(), &response)
            assert.NoError(t, err)
            assert.Equal(t, tt.expectedMsg, response["message"])

            // Verify mock expectations
            mockService.AssertExpectations(t)
        })
    }
}

func TestChangeAlumnusPassword(t *testing.T) {
    // Setup
    gin.SetMode(gin.TestMode)

    tests := []struct {
        name           string
        id             string
        requestBody    map[string]interface{}
        setupMock      func(*MockService)
        expectedStatus int
        expectedMsg    string
    }{
        {
            name: "Success",
            id:   "1",
            requestBody: map[string]interface{}{
                "password": "newpassword123",
            },
            setupMock: func(ms *MockService) {
                updatedAlumnus := &alumni.Alumni{
                    ID:       1,
                    MatricNo: "12345",
                    Name:     "Test User",
                }
                ms.On("UpdateAlumni", uint(1), mock.MatchedBy(func(a *alumni.Alumni) bool {
                    return a.Password != ""
                })).Return(updatedAlumnus, nil)
            },
            expectedStatus: http.StatusOK,
            expectedMsg:    "Password changed successfully",
        },
        {
            name: "Invalid ID",
            id:   "invalid",
            requestBody: map[string]interface{}{
                "password": "newpassword123",
            },
            setupMock:      func(ms *MockService) {},
            expectedStatus: http.StatusBadRequest,
            expectedMsg:    "Invalid alumnus ID",
        },
        {
            name: "Missing Password",
            id:   "1",
            requestBody: map[string]interface{}{
                "wrong_field": "value",
            },
            setupMock:      func(ms *MockService) {},
            expectedStatus: http.StatusBadRequest,
            expectedMsg:    "Invalid request body",
        },
        {
            name: "Empty Request Body",
            id:   "1",
            requestBody: map[string]interface{}{},
            setupMock:      func(ms *MockService) {},
            expectedStatus: http.StatusBadRequest,
            expectedMsg:    "Invalid request body",
        },
        {
            name: "Update Error",
            id:   "1",
            requestBody: map[string]interface{}{
                "password": "newpassword123",
            },
            setupMock: func(ms *MockService) {
                ms.On("UpdateAlumni", uint(1), mock.MatchedBy(func(a *alumni.Alumni) bool {
                    return a.Password != ""
                })).Return(nil, errors.New("update error"))
            },
            expectedStatus: http.StatusInternalServerError,
            expectedMsg:    "update error",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            // Create mock service and setup expectations
            mockService := new(MockService)
            if tt.setupMock != nil {
                tt.setupMock(mockService)
            }

            // Create handler using NewHandler constructor
            handler := alumni.NewHandler(mockService)

            // Create test context
            w := httptest.NewRecorder()
            c, _ := gin.CreateTestContext(w)

            // Set URL parameter
            c.Params = gin.Params{
                {Key: "id", Value: tt.id},
            }

            // Create request body
            jsonBody, _ := json.Marshal(tt.requestBody)
            c.Request = httptest.NewRequest("PUT", "/", bytes.NewBuffer(jsonBody))
            c.Request.Header.Set("Content-Type", "application/json")

            // Call the handler
            handler.ChangeAlumnusPassword(c)

            // Assert response
            assert.Equal(t, tt.expectedStatus, w.Code)

            var response map[string]interface{}
            err := json.Unmarshal(w.Body.Bytes(), &response)
            assert.NoError(t, err)
            assert.Equal(t, tt.expectedMsg, response["message"])

            // For successful case, verify the response data
            if tt.expectedStatus == http.StatusOK {
                assert.NotNil(t, response["data"])
                data := response["data"].(map[string]interface{})
                assert.Empty(t, data["password"], "Password should be empty in response")
            }

            // Verify mock expectations
            mockService.AssertExpectations(t)
        })
    }
}
