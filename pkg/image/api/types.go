package api

import (
	kapi "github.com/GoogleCloudPlatform/kubernetes/pkg/api"
	"github.com/fsouza/go-dockerclient"
)

// ImageList is a list of Image objects.
type ImageList struct {
	kapi.TypeMeta `json:",inline" yaml:",inline"`
	kapi.ListMeta `json:"metadata,omitempty" yaml:"metadata,omitempty"`
	Items         []Image `json:"items" yaml:"items"`
}

// Image is an immutable representation of a Docker image and metadata at a point in time.
type Image struct {
	kapi.TypeMeta        `json:",inline" yaml:",inline"`
	kapi.ObjectMeta      `json:"metadata,omitempty" yaml:"metadata,omitempty"`
	DockerImageReference string       `json:"dockerImageReference,omitempty" yaml:"dockerImageReference,omitempty"`
	DockerImageMetadata  docker.Image `json:"dockerImageMetadata,omitempty" yaml:"dockerImageMetadata,omitempty"`
}

// ImageRepositoryList is a list of ImageRepository objects.
type ImageRepositoryList struct {
	kapi.TypeMeta `json:",inline" yaml:",inline"`
	kapi.ListMeta `json:"metadata,omitempty" yaml:"metadata,omitempty"`
	Items         []ImageRepository `json:"items" yaml:"items"`
}

// ImageRepository stores a mapping of tags to images, metadata overrides that are applied
// when images are tagged in a repository, and an optional reference to a Docker image
// repository on a registry.
type ImageRepository struct {
	kapi.TypeMeta         `json:",inline" yaml:",inline"`
	kapi.ObjectMeta       `json:"metadata,omitempty" yaml:"metadata,omitempty"`
	DockerImageRepository string            `json:"dockerImageRepository,omitempty" yaml:"dockerImageRepository,omitempty"`
	Tags                  map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}

// TODO add metadata overrides

// ImageRepositoryMapping represents a mapping from a single tag to a Docker image as
// well as the reference to the Docker image repository the image came from.
type ImageRepositoryMapping struct {
	kapi.TypeMeta         `json:",inline" yaml:",inline"`
	kapi.ObjectMeta       `json:"metadata,omitempty" yaml:"metadata,omitempty"`
	DockerImageRepository string `json:"dockerImageRepository" yaml:"dockerImageRepository"`
	Image                 Image  `json:"image" yaml:"image"`
	Tag                   string `json:"tag" yaml:"tag"`
}
