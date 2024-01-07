package handler

import (
	"munayfund-api2/internal/core/domain"
	"munayfund-api2/internal/core/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

// Login realiza la autenticación de un usuario.
// @Summary Realiza la autenticación de un usuario.
// @Description Este endpoint permite a un usuario autenticarse, proporcionando el email y la contraseña.
// @ID loginUser
// @Tags users
// @Produce json
// @Accept json
// @Param loginUser body domain.LoginInput true "Credenciales de inicio de sesión del usuario"
// @Success 200 {object} gin.H "Token de autenticación generado"
// @Success 200 {object} domain.Project "Proyecto actualizado exitosamente"
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router /login [post]
func (h *UserHandler) Login(c *gin.Context) {
	var loginUser domain.User

	// Bind de los datos del cuerpo de la solicitud al struct User
	if err := c.ShouldBindJSON(&loginUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validación básica para asegurarse de que se proporcionen el email y la contraseña
	if loginUser.Email == "" || loginUser.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email y contraseña son obligatorios"})
		return
	}

	// Llamada al servicio de Login
	token, err := h.userService.Login(c.Request.Context(), &loginUser)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

// SignUp registra a un nuevo usuario en el sistema.
// @Summary Registra a un nuevo usuario en el sistema.
// @Description Este endpoint permite a un usuario registrarse proporcionando la información requerida, incluyendo email y contraseña.
// @ID signUpUser
// @Tags users
// @Produce json
// @Accept json
// @Param signUpUser body domain.User true "Información del nuevo usuario"
// @Success 200 {object} gin.H "Token de autenticación generado"
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router /signup [post]
func (h *UserHandler) SigUp(c *gin.Context) {
	var newUser domain.User

	// Bind de los datos del cuerpo de la solicitud al struct User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validación básica para asegurarse de que se proporcionen el email y la contraseña
	if newUser.Email == "" || newUser.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email y contraseña son obligatorios"})
		return
	}

	// Llamada al servicio de SignUp
	token, err := h.userService.Signup(c.Request.Context(), &newUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Devolver el token JWT al cliente
	c.JSON(http.StatusOK, gin.H{"token": token})
}

// Update actualiza la información del usuario.
// @Summary Actualiza la información del usuario.
// @Description Este endpoint permite a un usuario actualizar su información proporcionando los datos requeridos.
// @ID updateUser
// @Tags users
// @Security ApiKeyAuth
// @param Authorization header string true "Authorization"
// @Produce json
// @Accept json
// @Param id path string true "ID del usuario a actualizar"
// @Param updateUser body domain.User true "Información actualizada del usuario"
// @Success 200 {object} gin.H "Usuario actualizado"
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router /user/{id} [put]
func (h *UserHandler) Update(c *gin.Context) {
	var updateUser domain.User

	// Bind de los datos del cuerpo de la solicitud al struct User
	if err := c.ShouldBindJSON(&updateUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Obtener el ID del usuario de los parámetros de la ruta
	userID := c.Param("id")

	// Asignar el ID del usuario al struct updateUser
	updateUser.ID = userID

	// Llamada al servicio de Update
	updatedUser, err := h.userService.Update(c.Request.Context(), &updateUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newUser := struct {
		ID       string `json:"id" bson:"_id"`
		Username string `json:"username" bson:"username"`
		WalletID string `json:"wallet_id" bson:"wallet_id"`
		Email    string `json:"email" bson:"email"`
	}{
		ID:       updatedUser.ID,
		Username: updatedUser.Username,
		WalletID: updatedUser.WalletID,
		Email:    updatedUser.Email,
	}

	// Devolver el usuario actualizado al cliente
	c.JSON(http.StatusOK, newUser)
}

// Delete elimina un usuario existente.
// @Summary Elimina un usuario existente.
// @Description Este endpoint permite a un usuario eliminar su cuenta proporcionando su ID.
// @ID deleteUser
// @Tags users
// @Security ApiKeyAuth
// @param Authorization header string true "Authorization"
// @Produce json
// @Param id path string true "ID del usuario a eliminar"
// @Success 200 {object} gin.H "Respuesta de éxito al eliminar el usuario"
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router /user/{id} [delete]
func (h *UserHandler) Delete(c *gin.Context) {
	// Obtener el ID del usuario de los parámetros de la ruta
	userID := c.Param("id")

	// Llamada al servicio de Delete
	err := h.userService.Delete(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Devolver respuesta de éxito al cliente
	c.JSON(http.StatusOK, gin.H{"message": "Successfully deleted"})
}
