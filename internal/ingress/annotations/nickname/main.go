package nickname

import ( 
	networking "k8s.io/api/networking/v1beta1"

	"k8s.io/ingress-nginx/internal/ingress/annotations/parser"
	"k8s.io/ingress-nginx/internal/ingress/resolver"
)

type nickname struct { 
	r resolver.Resolver
}

// NewParser ...
func NewParser(r resolver.Resolver) parser.IngressAnnotation { 
	return nickname{r}
}

func (a nickname) Parse(ing *networking.Ingress) (interface{}, error) { 
	return parser.GetStringAnnotation("nickname", ing)
}