// moq -pkg todo_test -out ./test/todo_mock.go . ITodoUsecase ITodoService ITodoRepository
// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package todo_test

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go-todo-app/feature/todo"
	"go-todo-app/feature/user"
	"gorm.io/gorm"
	"sync"
)

// Ensure, that ITodoUsecaseMock does implement todo.ITodoUsecase.
// If this is not the case, regenerate this file with moq.
var _ todo.ITodoUsecase = &ITodoUsecaseMock{}

// ITodoUsecaseMock is a mock implementation of todo.ITodoUsecase.
//
//	func TestSomethingThatUsesITodoUsecase(t *testing.T) {
//
//		// make and configure a mocked todo.ITodoUsecase
//		mockedITodoUsecase := &ITodoUsecaseMock{
//			CreateFunc: func(ctx *gin.Context, userContext user.UserContext, req todo.CreateTodoRequest)  {
//				panic("mock out the Create method")
//			},
//			DeleteFunc: func(ctx *gin.Context, userContext user.UserContext, id uuid.UUID)  {
//				panic("mock out the Delete method")
//			},
//			FindAllFunc: func(ctx *gin.Context, userContext user.UserContext) []todo.Todo {
//				panic("mock out the FindAll method")
//			},
//			FindByIdFunc: func(ctx *gin.Context, userContext user.UserContext, id uuid.UUID) todo.Todo {
//				panic("mock out the FindById method")
//			},
//		}
//
//		// use mockedITodoUsecase in code that requires todo.ITodoUsecase
//		// and then make assertions.
//
//	}
type ITodoUsecaseMock struct {
	// CreateFunc mocks the Create method.
	CreateFunc func(ctx *gin.Context, userContext user.UserContext, req todo.CreateTodoRequest)

	// DeleteFunc mocks the Delete method.
	DeleteFunc func(ctx *gin.Context, userContext user.UserContext, id uuid.UUID)

	// FindAllFunc mocks the FindAll method.
	FindAllFunc func(ctx *gin.Context, userContext user.UserContext) []todo.Todo

	// FindByIdFunc mocks the FindById method.
	FindByIdFunc func(ctx *gin.Context, userContext user.UserContext, id uuid.UUID) todo.Todo

	// calls tracks calls to the methods.
	calls struct {
		// Create holds details about calls to the Create method.
		Create []struct {
			// Ctx is the ctx argument value.
			Ctx *gin.Context
			// UserContext is the userContext argument value.
			UserContext user.UserContext
			// Req is the req argument value.
			Req todo.CreateTodoRequest
		}
		// Delete holds details about calls to the Delete method.
		Delete []struct {
			// Ctx is the ctx argument value.
			Ctx *gin.Context
			// UserContext is the userContext argument value.
			UserContext user.UserContext
			// ID is the id argument value.
			ID uuid.UUID
		}
		// FindAll holds details about calls to the FindAll method.
		FindAll []struct {
			// Ctx is the ctx argument value.
			Ctx *gin.Context
			// UserContext is the userContext argument value.
			UserContext user.UserContext
		}
		// FindById holds details about calls to the FindById method.
		FindById []struct {
			// Ctx is the ctx argument value.
			Ctx *gin.Context
			// UserContext is the userContext argument value.
			UserContext user.UserContext
			// ID is the id argument value.
			ID uuid.UUID
		}
	}
	lockCreate   sync.RWMutex
	lockDelete   sync.RWMutex
	lockFindAll  sync.RWMutex
	lockFindById sync.RWMutex
}

// Create calls CreateFunc.
func (mock *ITodoUsecaseMock) Create(ctx *gin.Context, userContext user.UserContext, req todo.CreateTodoRequest) {
	if mock.CreateFunc == nil {
		panic("ITodoUsecaseMock.CreateFunc: method is nil but ITodoUsecase.Create was just called")
	}
	callInfo := struct {
		Ctx         *gin.Context
		UserContext user.UserContext
		Req         todo.CreateTodoRequest
	}{
		Ctx:         ctx,
		UserContext: userContext,
		Req:         req,
	}
	mock.lockCreate.Lock()
	mock.calls.Create = append(mock.calls.Create, callInfo)
	mock.lockCreate.Unlock()
	mock.CreateFunc(ctx, userContext, req)
}

// CreateCalls gets all the calls that were made to Create.
// Check the length with:
//
//	len(mockedITodoUsecase.CreateCalls())
func (mock *ITodoUsecaseMock) CreateCalls() []struct {
	Ctx         *gin.Context
	UserContext user.UserContext
	Req         todo.CreateTodoRequest
} {
	var calls []struct {
		Ctx         *gin.Context
		UserContext user.UserContext
		Req         todo.CreateTodoRequest
	}
	mock.lockCreate.RLock()
	calls = mock.calls.Create
	mock.lockCreate.RUnlock()
	return calls
}

// Delete calls DeleteFunc.
func (mock *ITodoUsecaseMock) Delete(ctx *gin.Context, userContext user.UserContext, id uuid.UUID) {
	if mock.DeleteFunc == nil {
		panic("ITodoUsecaseMock.DeleteFunc: method is nil but ITodoUsecase.Delete was just called")
	}
	callInfo := struct {
		Ctx         *gin.Context
		UserContext user.UserContext
		ID          uuid.UUID
	}{
		Ctx:         ctx,
		UserContext: userContext,
		ID:          id,
	}
	mock.lockDelete.Lock()
	mock.calls.Delete = append(mock.calls.Delete, callInfo)
	mock.lockDelete.Unlock()
	mock.DeleteFunc(ctx, userContext, id)
}

// DeleteCalls gets all the calls that were made to Delete.
// Check the length with:
//
//	len(mockedITodoUsecase.DeleteCalls())
func (mock *ITodoUsecaseMock) DeleteCalls() []struct {
	Ctx         *gin.Context
	UserContext user.UserContext
	ID          uuid.UUID
} {
	var calls []struct {
		Ctx         *gin.Context
		UserContext user.UserContext
		ID          uuid.UUID
	}
	mock.lockDelete.RLock()
	calls = mock.calls.Delete
	mock.lockDelete.RUnlock()
	return calls
}

// FindAll calls FindAllFunc.
func (mock *ITodoUsecaseMock) FindAll(ctx *gin.Context, userContext user.UserContext) []todo.Todo {
	if mock.FindAllFunc == nil {
		panic("ITodoUsecaseMock.FindAllFunc: method is nil but ITodoUsecase.FindAll was just called")
	}
	callInfo := struct {
		Ctx         *gin.Context
		UserContext user.UserContext
	}{
		Ctx:         ctx,
		UserContext: userContext,
	}
	mock.lockFindAll.Lock()
	mock.calls.FindAll = append(mock.calls.FindAll, callInfo)
	mock.lockFindAll.Unlock()
	return mock.FindAllFunc(ctx, userContext)
}

// FindAllCalls gets all the calls that were made to FindAll.
// Check the length with:
//
//	len(mockedITodoUsecase.FindAllCalls())
func (mock *ITodoUsecaseMock) FindAllCalls() []struct {
	Ctx         *gin.Context
	UserContext user.UserContext
} {
	var calls []struct {
		Ctx         *gin.Context
		UserContext user.UserContext
	}
	mock.lockFindAll.RLock()
	calls = mock.calls.FindAll
	mock.lockFindAll.RUnlock()
	return calls
}

// FindById calls FindByIdFunc.
func (mock *ITodoUsecaseMock) FindById(ctx *gin.Context, userContext user.UserContext, id uuid.UUID) todo.Todo {
	if mock.FindByIdFunc == nil {
		panic("ITodoUsecaseMock.FindByIdFunc: method is nil but ITodoUsecase.FindById was just called")
	}
	callInfo := struct {
		Ctx         *gin.Context
		UserContext user.UserContext
		ID          uuid.UUID
	}{
		Ctx:         ctx,
		UserContext: userContext,
		ID:          id,
	}
	mock.lockFindById.Lock()
	mock.calls.FindById = append(mock.calls.FindById, callInfo)
	mock.lockFindById.Unlock()
	return mock.FindByIdFunc(ctx, userContext, id)
}

// FindByIdCalls gets all the calls that were made to FindById.
// Check the length with:
//
//	len(mockedITodoUsecase.FindByIdCalls())
func (mock *ITodoUsecaseMock) FindByIdCalls() []struct {
	Ctx         *gin.Context
	UserContext user.UserContext
	ID          uuid.UUID
} {
	var calls []struct {
		Ctx         *gin.Context
		UserContext user.UserContext
		ID          uuid.UUID
	}
	mock.lockFindById.RLock()
	calls = mock.calls.FindById
	mock.lockFindById.RUnlock()
	return calls
}

// Ensure, that ITodoServiceMock does implement todo.ITodoService.
// If this is not the case, regenerate this file with moq.
var _ todo.ITodoService = &ITodoServiceMock{}

// ITodoServiceMock is a mock implementation of todo.ITodoService.
//
//	func TestSomethingThatUsesITodoService(t *testing.T) {
//
//		// make and configure a mocked todo.ITodoService
//		mockedITodoService := &ITodoServiceMock{
//			CreateFunc: func(ctx *gin.Context, userContext user.UserContext, t *todo.Todo, session *gorm.DB)  {
//				panic("mock out the Create method")
//			},
//			DeleteFunc: func(ctx *gin.Context, userContext user.UserContext, id uuid.UUID, session *gorm.DB)  {
//				panic("mock out the Delete method")
//			},
//			FindAllFunc: func(ctx *gin.Context, userContext user.UserContext, session *gorm.DB) []todo.Todo {
//				panic("mock out the FindAll method")
//			},
//			FindByIdFunc: func(ctx *gin.Context, userContext user.UserContext, id uuid.UUID, session *gorm.DB) todo.Todo {
//				panic("mock out the FindById method")
//			},
//		}
//
//		// use mockedITodoService in code that requires todo.ITodoService
//		// and then make assertions.
//
//	}
type ITodoServiceMock struct {
	// CreateFunc mocks the Create method.
	CreateFunc func(ctx *gin.Context, userContext user.UserContext, t *todo.Todo, session *gorm.DB)

	// DeleteFunc mocks the Delete method.
	DeleteFunc func(ctx *gin.Context, userContext user.UserContext, id uuid.UUID, session *gorm.DB)

	// FindAllFunc mocks the FindAll method.
	FindAllFunc func(ctx *gin.Context, userContext user.UserContext, session *gorm.DB) []todo.Todo

	// FindByIdFunc mocks the FindById method.
	FindByIdFunc func(ctx *gin.Context, userContext user.UserContext, id uuid.UUID, session *gorm.DB) todo.Todo

	// calls tracks calls to the methods.
	calls struct {
		// Create holds details about calls to the Create method.
		Create []struct {
			// Ctx is the ctx argument value.
			Ctx *gin.Context
			// UserContext is the userContext argument value.
			UserContext user.UserContext
			// T is the t argument value.
			T *todo.Todo
			// Session is the session argument value.
			Session *gorm.DB
		}
		// Delete holds details about calls to the Delete method.
		Delete []struct {
			// Ctx is the ctx argument value.
			Ctx *gin.Context
			// UserContext is the userContext argument value.
			UserContext user.UserContext
			// ID is the id argument value.
			ID uuid.UUID
			// Session is the session argument value.
			Session *gorm.DB
		}
		// FindAll holds details about calls to the FindAll method.
		FindAll []struct {
			// Ctx is the ctx argument value.
			Ctx *gin.Context
			// UserContext is the userContext argument value.
			UserContext user.UserContext
			// Session is the session argument value.
			Session *gorm.DB
		}
		// FindById holds details about calls to the FindById method.
		FindById []struct {
			// Ctx is the ctx argument value.
			Ctx *gin.Context
			// UserContext is the userContext argument value.
			UserContext user.UserContext
			// ID is the id argument value.
			ID uuid.UUID
			// Session is the session argument value.
			Session *gorm.DB
		}
	}
	lockCreate   sync.RWMutex
	lockDelete   sync.RWMutex
	lockFindAll  sync.RWMutex
	lockFindById sync.RWMutex
}

// Create calls CreateFunc.
func (mock *ITodoServiceMock) Create(ctx *gin.Context, userContext user.UserContext, t *todo.Todo, session *gorm.DB) {
	if mock.CreateFunc == nil {
		panic("ITodoServiceMock.CreateFunc: method is nil but ITodoService.Create was just called")
	}
	callInfo := struct {
		Ctx         *gin.Context
		UserContext user.UserContext
		T           *todo.Todo
		Session     *gorm.DB
	}{
		Ctx:         ctx,
		UserContext: userContext,
		T:           t,
		Session:     session,
	}
	mock.lockCreate.Lock()
	mock.calls.Create = append(mock.calls.Create, callInfo)
	mock.lockCreate.Unlock()
	mock.CreateFunc(ctx, userContext, t, session)
}

// CreateCalls gets all the calls that were made to Create.
// Check the length with:
//
//	len(mockedITodoService.CreateCalls())
func (mock *ITodoServiceMock) CreateCalls() []struct {
	Ctx         *gin.Context
	UserContext user.UserContext
	T           *todo.Todo
	Session     *gorm.DB
} {
	var calls []struct {
		Ctx         *gin.Context
		UserContext user.UserContext
		T           *todo.Todo
		Session     *gorm.DB
	}
	mock.lockCreate.RLock()
	calls = mock.calls.Create
	mock.lockCreate.RUnlock()
	return calls
}

// Delete calls DeleteFunc.
func (mock *ITodoServiceMock) Delete(ctx *gin.Context, userContext user.UserContext, id uuid.UUID, session *gorm.DB) {
	if mock.DeleteFunc == nil {
		panic("ITodoServiceMock.DeleteFunc: method is nil but ITodoService.Delete was just called")
	}
	callInfo := struct {
		Ctx         *gin.Context
		UserContext user.UserContext
		ID          uuid.UUID
		Session     *gorm.DB
	}{
		Ctx:         ctx,
		UserContext: userContext,
		ID:          id,
		Session:     session,
	}
	mock.lockDelete.Lock()
	mock.calls.Delete = append(mock.calls.Delete, callInfo)
	mock.lockDelete.Unlock()
	mock.DeleteFunc(ctx, userContext, id, session)
}

// DeleteCalls gets all the calls that were made to Delete.
// Check the length with:
//
//	len(mockedITodoService.DeleteCalls())
func (mock *ITodoServiceMock) DeleteCalls() []struct {
	Ctx         *gin.Context
	UserContext user.UserContext
	ID          uuid.UUID
	Session     *gorm.DB
} {
	var calls []struct {
		Ctx         *gin.Context
		UserContext user.UserContext
		ID          uuid.UUID
		Session     *gorm.DB
	}
	mock.lockDelete.RLock()
	calls = mock.calls.Delete
	mock.lockDelete.RUnlock()
	return calls
}

// FindAll calls FindAllFunc.
func (mock *ITodoServiceMock) FindAll(ctx *gin.Context, userContext user.UserContext, session *gorm.DB) []todo.Todo {
	if mock.FindAllFunc == nil {
		panic("ITodoServiceMock.FindAllFunc: method is nil but ITodoService.FindAll was just called")
	}
	callInfo := struct {
		Ctx         *gin.Context
		UserContext user.UserContext
		Session     *gorm.DB
	}{
		Ctx:         ctx,
		UserContext: userContext,
		Session:     session,
	}
	mock.lockFindAll.Lock()
	mock.calls.FindAll = append(mock.calls.FindAll, callInfo)
	mock.lockFindAll.Unlock()
	return mock.FindAllFunc(ctx, userContext, session)
}

// FindAllCalls gets all the calls that were made to FindAll.
// Check the length with:
//
//	len(mockedITodoService.FindAllCalls())
func (mock *ITodoServiceMock) FindAllCalls() []struct {
	Ctx         *gin.Context
	UserContext user.UserContext
	Session     *gorm.DB
} {
	var calls []struct {
		Ctx         *gin.Context
		UserContext user.UserContext
		Session     *gorm.DB
	}
	mock.lockFindAll.RLock()
	calls = mock.calls.FindAll
	mock.lockFindAll.RUnlock()
	return calls
}

// FindById calls FindByIdFunc.
func (mock *ITodoServiceMock) FindById(ctx *gin.Context, userContext user.UserContext, id uuid.UUID, session *gorm.DB) todo.Todo {
	if mock.FindByIdFunc == nil {
		panic("ITodoServiceMock.FindByIdFunc: method is nil but ITodoService.FindById was just called")
	}
	callInfo := struct {
		Ctx         *gin.Context
		UserContext user.UserContext
		ID          uuid.UUID
		Session     *gorm.DB
	}{
		Ctx:         ctx,
		UserContext: userContext,
		ID:          id,
		Session:     session,
	}
	mock.lockFindById.Lock()
	mock.calls.FindById = append(mock.calls.FindById, callInfo)
	mock.lockFindById.Unlock()
	return mock.FindByIdFunc(ctx, userContext, id, session)
}

// FindByIdCalls gets all the calls that were made to FindById.
// Check the length with:
//
//	len(mockedITodoService.FindByIdCalls())
func (mock *ITodoServiceMock) FindByIdCalls() []struct {
	Ctx         *gin.Context
	UserContext user.UserContext
	ID          uuid.UUID
	Session     *gorm.DB
} {
	var calls []struct {
		Ctx         *gin.Context
		UserContext user.UserContext
		ID          uuid.UUID
		Session     *gorm.DB
	}
	mock.lockFindById.RLock()
	calls = mock.calls.FindById
	mock.lockFindById.RUnlock()
	return calls
}

// Ensure, that ITodoRepositoryMock does implement todo.ITodoRepository.
// If this is not the case, regenerate this file with moq.
var _ todo.ITodoRepository = &ITodoRepositoryMock{}

// ITodoRepositoryMock is a mock implementation of todo.ITodoRepository.
//
//	func TestSomethingThatUsesITodoRepository(t *testing.T) {
//
//		// make and configure a mocked todo.ITodoRepository
//		mockedITodoRepository := &ITodoRepositoryMock{
//			CreateFunc: func(ctx *gin.Context, userContext user.UserContext, t *todo.Todo, session *gorm.DB)  {
//				panic("mock out the Create method")
//			},
//			DeleteFunc: func(ctx *gin.Context, userContext user.UserContext, id uuid.UUID, session *gorm.DB)  {
//				panic("mock out the Delete method")
//			},
//			FindAllFunc: func(ctx *gin.Context, userContext user.UserContext, session *gorm.DB) []todo.Todo {
//				panic("mock out the FindAll method")
//			},
//			FindByIdFunc: func(ctx *gin.Context, userContext user.UserContext, id uuid.UUID, session *gorm.DB) todo.Todo {
//				panic("mock out the FindById method")
//			},
//		}
//
//		// use mockedITodoRepository in code that requires todo.ITodoRepository
//		// and then make assertions.
//
//	}
type ITodoRepositoryMock struct {
	// CreateFunc mocks the Create method.
	CreateFunc func(ctx *gin.Context, userContext user.UserContext, t *todo.Todo, session *gorm.DB)

	// DeleteFunc mocks the Delete method.
	DeleteFunc func(ctx *gin.Context, userContext user.UserContext, id uuid.UUID, session *gorm.DB)

	// FindAllFunc mocks the FindAll method.
	FindAllFunc func(ctx *gin.Context, userContext user.UserContext, session *gorm.DB) []todo.Todo

	// FindByIdFunc mocks the FindById method.
	FindByIdFunc func(ctx *gin.Context, userContext user.UserContext, id uuid.UUID, session *gorm.DB) todo.Todo

	// calls tracks calls to the methods.
	calls struct {
		// Create holds details about calls to the Create method.
		Create []struct {
			// Ctx is the ctx argument value.
			Ctx *gin.Context
			// UserContext is the userContext argument value.
			UserContext user.UserContext
			// T is the t argument value.
			T *todo.Todo
			// Session is the session argument value.
			Session *gorm.DB
		}
		// Delete holds details about calls to the Delete method.
		Delete []struct {
			// Ctx is the ctx argument value.
			Ctx *gin.Context
			// UserContext is the userContext argument value.
			UserContext user.UserContext
			// ID is the id argument value.
			ID uuid.UUID
			// Session is the session argument value.
			Session *gorm.DB
		}
		// FindAll holds details about calls to the FindAll method.
		FindAll []struct {
			// Ctx is the ctx argument value.
			Ctx *gin.Context
			// UserContext is the userContext argument value.
			UserContext user.UserContext
			// Session is the session argument value.
			Session *gorm.DB
		}
		// FindById holds details about calls to the FindById method.
		FindById []struct {
			// Ctx is the ctx argument value.
			Ctx *gin.Context
			// UserContext is the userContext argument value.
			UserContext user.UserContext
			// ID is the id argument value.
			ID uuid.UUID
			// Session is the session argument value.
			Session *gorm.DB
		}
	}
	lockCreate   sync.RWMutex
	lockDelete   sync.RWMutex
	lockFindAll  sync.RWMutex
	lockFindById sync.RWMutex
}

// Create calls CreateFunc.
func (mock *ITodoRepositoryMock) Create(ctx *gin.Context, userContext user.UserContext, t *todo.Todo, session *gorm.DB) {
	if mock.CreateFunc == nil {
		panic("ITodoRepositoryMock.CreateFunc: method is nil but ITodoRepository.Create was just called")
	}
	callInfo := struct {
		Ctx         *gin.Context
		UserContext user.UserContext
		T           *todo.Todo
		Session     *gorm.DB
	}{
		Ctx:         ctx,
		UserContext: userContext,
		T:           t,
		Session:     session,
	}
	mock.lockCreate.Lock()
	mock.calls.Create = append(mock.calls.Create, callInfo)
	mock.lockCreate.Unlock()
	mock.CreateFunc(ctx, userContext, t, session)
}

// CreateCalls gets all the calls that were made to Create.
// Check the length with:
//
//	len(mockedITodoRepository.CreateCalls())
func (mock *ITodoRepositoryMock) CreateCalls() []struct {
	Ctx         *gin.Context
	UserContext user.UserContext
	T           *todo.Todo
	Session     *gorm.DB
} {
	var calls []struct {
		Ctx         *gin.Context
		UserContext user.UserContext
		T           *todo.Todo
		Session     *gorm.DB
	}
	mock.lockCreate.RLock()
	calls = mock.calls.Create
	mock.lockCreate.RUnlock()
	return calls
}

// Delete calls DeleteFunc.
func (mock *ITodoRepositoryMock) Delete(ctx *gin.Context, userContext user.UserContext, id uuid.UUID, session *gorm.DB) {
	if mock.DeleteFunc == nil {
		panic("ITodoRepositoryMock.DeleteFunc: method is nil but ITodoRepository.Delete was just called")
	}
	callInfo := struct {
		Ctx         *gin.Context
		UserContext user.UserContext
		ID          uuid.UUID
		Session     *gorm.DB
	}{
		Ctx:         ctx,
		UserContext: userContext,
		ID:          id,
		Session:     session,
	}
	mock.lockDelete.Lock()
	mock.calls.Delete = append(mock.calls.Delete, callInfo)
	mock.lockDelete.Unlock()
	mock.DeleteFunc(ctx, userContext, id, session)
}

// DeleteCalls gets all the calls that were made to Delete.
// Check the length with:
//
//	len(mockedITodoRepository.DeleteCalls())
func (mock *ITodoRepositoryMock) DeleteCalls() []struct {
	Ctx         *gin.Context
	UserContext user.UserContext
	ID          uuid.UUID
	Session     *gorm.DB
} {
	var calls []struct {
		Ctx         *gin.Context
		UserContext user.UserContext
		ID          uuid.UUID
		Session     *gorm.DB
	}
	mock.lockDelete.RLock()
	calls = mock.calls.Delete
	mock.lockDelete.RUnlock()
	return calls
}

// FindAll calls FindAllFunc.
func (mock *ITodoRepositoryMock) FindAll(ctx *gin.Context, userContext user.UserContext, session *gorm.DB) []todo.Todo {
	if mock.FindAllFunc == nil {
		panic("ITodoRepositoryMock.FindAllFunc: method is nil but ITodoRepository.FindAll was just called")
	}
	callInfo := struct {
		Ctx         *gin.Context
		UserContext user.UserContext
		Session     *gorm.DB
	}{
		Ctx:         ctx,
		UserContext: userContext,
		Session:     session,
	}
	mock.lockFindAll.Lock()
	mock.calls.FindAll = append(mock.calls.FindAll, callInfo)
	mock.lockFindAll.Unlock()
	return mock.FindAllFunc(ctx, userContext, session)
}

// FindAllCalls gets all the calls that were made to FindAll.
// Check the length with:
//
//	len(mockedITodoRepository.FindAllCalls())
func (mock *ITodoRepositoryMock) FindAllCalls() []struct {
	Ctx         *gin.Context
	UserContext user.UserContext
	Session     *gorm.DB
} {
	var calls []struct {
		Ctx         *gin.Context
		UserContext user.UserContext
		Session     *gorm.DB
	}
	mock.lockFindAll.RLock()
	calls = mock.calls.FindAll
	mock.lockFindAll.RUnlock()
	return calls
}

// FindById calls FindByIdFunc.
func (mock *ITodoRepositoryMock) FindById(ctx *gin.Context, userContext user.UserContext, id uuid.UUID, session *gorm.DB) todo.Todo {
	if mock.FindByIdFunc == nil {
		panic("ITodoRepositoryMock.FindByIdFunc: method is nil but ITodoRepository.FindById was just called")
	}
	callInfo := struct {
		Ctx         *gin.Context
		UserContext user.UserContext
		ID          uuid.UUID
		Session     *gorm.DB
	}{
		Ctx:         ctx,
		UserContext: userContext,
		ID:          id,
		Session:     session,
	}
	mock.lockFindById.Lock()
	mock.calls.FindById = append(mock.calls.FindById, callInfo)
	mock.lockFindById.Unlock()
	return mock.FindByIdFunc(ctx, userContext, id, session)
}

// FindByIdCalls gets all the calls that were made to FindById.
// Check the length with:
//
//	len(mockedITodoRepository.FindByIdCalls())
func (mock *ITodoRepositoryMock) FindByIdCalls() []struct {
	Ctx         *gin.Context
	UserContext user.UserContext
	ID          uuid.UUID
	Session     *gorm.DB
} {
	var calls []struct {
		Ctx         *gin.Context
		UserContext user.UserContext
		ID          uuid.UUID
		Session     *gorm.DB
	}
	mock.lockFindById.RLock()
	calls = mock.calls.FindById
	mock.lockFindById.RUnlock()
	return calls
}
