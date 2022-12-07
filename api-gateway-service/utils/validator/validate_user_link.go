package validator

import (
	"api-gateway-service/internal/pb/users_pb"
	"errors"
	"regexp"
)

var urlRegexpVK *regexp.Regexp
var urlRegexpSoundCloud *regexp.Regexp
var urlRegexpDefault *regexp.Regexp

func init() {
	urlRegexpVK = regexp.MustCompile(`^(https://)?vk\.com/[a-zA-Z0-9]+$`)
	urlRegexpSoundCloud = regexp.MustCompile(`^(https://)?(on\.|m\.)?soundcloud\.com/[a-zA-Z0-9]+$`)
	urlRegexpDefault = regexp.MustCompile(`^(https://)?[a-zA-Z0-9-_]+\.[a-zA-Z0-9-_/.]+$`)
}

func ValidateLinkType(linkType users_pb.LinkType) error {
	switch linkType {
	case users_pb.LinkType_VK, users_pb.LinkType_SOUNDCLOUD, users_pb.LinkType_CUSTOM_WEBSITE:
		return nil
	}
	return errors.New("validate user link: invalid link type")
}

func ValidateURL(URL string, linkType users_pb.LinkType) error {
	if urlRegexpVK == nil {
		urlRegexpVK = regexp.MustCompile(`^(https://)?vk\.com/[a-zA-Z0-9]+$`)
	}
	if urlRegexpSoundCloud == nil {
		urlRegexpSoundCloud = regexp.MustCompile(`^(https://)?(on\.|m\.)?soundcloud\.com/[a-zA-Z0-9]+$`)
	}
	if urlRegexpDefault == nil {
		urlRegexpSoundCloud = regexp.MustCompile(`^(https://)?(on\.|m\.)?soundcloud\.com/[a-zA-Z0-9]+$`)
	}
	switch linkType {
	case users_pb.LinkType_VK:
		if !urlRegexpVK.Match([]byte(URL)) {
			return errors.New("validate user link: invalid url")
		}
	case users_pb.LinkType_SOUNDCLOUD:
		if !urlRegexpSoundCloud.Match([]byte(URL)) {
			return errors.New("validate user link: invalid url")
		}
	default:
		if !urlRegexpDefault.Match([]byte(URL)) {
			return errors.New("validate user link: invalid url")
		}
	}
	return nil
}

func ValidateUserLink(userLink *users_pb.UserLink) error {
	if len(userLink.Url) == 0 {
		return errors.New("validate user link: url is empty")
	}
	err := ValidateLinkType(userLink.Type)
	if err != nil {
		return err
	}

	return ValidateURL(userLink.Url, userLink.Type)
}
