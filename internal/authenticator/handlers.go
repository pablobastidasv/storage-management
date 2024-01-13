package authenticator

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/gob"
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/session"
)

type Handlers struct {
	store         *session.Store
	authenticator *Authenticator
}

func NewHandlers(store *session.Store, authenticator *Authenticator) *Handlers {
	// To store custom types in our cookies,
	// we must first register them using gob.Register
	gob.Register(map[string]interface{}{})

	return &Handlers{
		store:         store,
		authenticator: authenticator,
	}
}

func (h *Handlers) HandleLogin(c *fiber.Ctx) error {
	session, err := h.store.Get(c)
	if err != nil {
		return err
	}

	state, err := generateRandomState()
	if err != nil {
		return err
	}

	session.Set("state", state)
	if err := session.Save(); err != nil {
		return err
	}

	return c.Redirect(h.authenticator.AuthCodeURL(state), http.StatusTemporaryRedirect)
}

func generateRandomState() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	state := base64.StdEncoding.EncodeToString(b)
	return state, nil
}

func (h *Handlers) HandleCallback(c *fiber.Ctx) error {
	session, err := h.store.Get(c)
	if err != nil {
		return err
	}

	// Exchange an authorization code for a token
	token, err := h.authenticator.Exchange(c.Context(), c.Query("code"))
	if err != nil {
		return err
	}

	idToken, err := h.authenticator.VerifyIDToken(c.Context(), token)
	if err != nil {
		return err
	}

	var profile map[string]interface{}
	if err := idToken.Claims(&profile); err != nil {
		return err
	}

	session.Set("access_token", token.AccessToken)
	session.Set("profile", profile)
	if err := session.Save(); err != nil {
		return err
	}

	return c.Redirect("/", http.StatusTemporaryRedirect)
}

func (h *Handlers) HandleLogout(c *fiber.Ctx) error {
	issuer := os.Getenv("AUTH_ISSUER")
	logoutUrl, err := url.Parse(issuer)
	if err != nil {
		return err
	}

	schema := c.Protocol()

	returnTo, err := url.Parse(fmt.Sprintf("%s://%s", schema, c.Request().Host()))
	if err != nil {
		return err
	}

	parameters := url.Values{}
	parameters.Add("returnTo", returnTo.String())
	parameters.Add("client_id", h.authenticator.ClientID)
	logoutUrl.RawQuery = parameters.Encode()

	return c.Redirect(logoutUrl.String(), http.StatusTemporaryRedirect)
}

func (h *Handlers) IsAuthenticated(c *fiber.Ctx) error {
	theSession, err := h.store.Get(c)
	if err != nil {
		return err
	}

	if theSession.Get("profile") == nil {
		return c.Redirect("/login", http.StatusSeeOther)
	} else {
		return c.Next()
	}
}

func (h *Handlers) HandleGetUsersMe(c *fiber.Ctx) error {
	theSession, err := h.store.Get(c)
	if err != nil {
		log.Errorf("error loading session, server says %v\n", err)

		return c.SendString("<div>Error loading session</div>")
	}

	profile := theSession.Get("profile")

	// Send the fragment with user info
	return c.Render("profile", profile)
}
