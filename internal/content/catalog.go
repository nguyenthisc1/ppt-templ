package content

import "strings"

type Listing struct {
	Slug        string
	Title       string
	Category    string
	Location    string
	PriceLabel  string
	Beds        int
	Baths       int
	Area        string
	Badge       string
	Image       string
	Summary     string
	Gallery     []string
	Headline    string
	Description string
	Options     []string
	Info        []DetailItem
	Features    []string
}

type DetailItem struct {
	Label string
	Value string
}

type Project struct {
	Title      string
	Status     string
	Location   string
	Delivery   string
	PriceLabel string
	Image      string
	Summary    string
}

type Article struct {
	Title     string
	Category  string
	Published string
	Image     string
	Summary   string
}

func Listings() []Listing {
	return []Listing{
		{Slug: "the-crest-penthouse", Title: "The Crest Penthouse", Category: "sale", Location: "Thu Thiem, Thu Duc", PriceLabel: "$1.45M", Beds: 4, Baths: 3, Area: "278 sqm", Badge: "Skyline", Image: "/assets/img/listing_1.avif", Summary: "A top-floor residence with panoramic skyline views, private lift access, and a deep entertaining terrace.", Gallery: []string{"/assets/img/listing_1.avif", "/assets/img/listing_2.avif", "/assets/img/listing_3.avif"}, Headline: "A skyline penthouse designed for buyers who want privacy, scale, and dramatic city views.", Description: "The Crest Penthouse combines oversized entertaining space, a refined arrival sequence, and an indoor-outdoor living rhythm that works for both family life and hosted evenings.", Options: []string{"Private lift lobby", "Corner living room with full-height glazing", "Entertaining terrace with skyline outlook", "Dual-kitchen arrangement for formal and daily use"}, Info: []DetailItem{{Label: "Property type", Value: "Penthouse residence"}, {Label: "Availability", Value: "Ready to view"}, {Label: "Parking", Value: "2 private bays"}, {Label: "Furnishing", Value: "Designer fit-out"}}, Features: []string{"Calm material palette with marble and oak detailing", "Separate family lounge and guest-facing salon", "Primary suite with dressing room and spa bath", "Immediate access to riverside retail and wellness amenities"}},
		{Slug: "riverlight-residences", Title: "Riverlight Residences", Category: "sale", Location: "District 2, Ho Chi Minh City", PriceLabel: "$920K", Beds: 3, Baths: 2, Area: "162 sqm", Badge: "Waterfront", Image: "/assets/img/listing_2.avif", Summary: "Open-plan living with a marble kitchen, wraparound glazing, and direct access to a landscaped riverwalk.", Gallery: []string{"/assets/img/listing_2.avif", "/assets/img/listing_1.avif", "/assets/img/listing_10.avif"}, Headline: "Waterfront family living with a more relaxed, everyday luxury mood.", Description: "Riverlight Residences balances open entertaining with practical family planning, offering wide living zones, river-facing glazing, and a strong lifestyle location.", Options: []string{"3-bedroom family plan", "Riverwalk access", "Clubhouse and pool deck", "Walkable retail and dining"}, Info: []DetailItem{{Label: "Property type", Value: "Apartment"}, {Label: "Availability", Value: "Ready to move"}, {Label: "Parking", Value: "1 basement bay"}, {Label: "Furnishing", Value: "Semi-furnished"}}, Features: []string{"Wide living-dining frontage with river outlook", "Primary bedroom separated from secondary rooms", "Amenity floor with pool, lounge, and fitness spaces", "Well-suited for owner-occupiers upgrading from central districts"}},
		{Slug: "lumiere-garden-villa", Title: "Lumiere Garden Villa", Category: "sale", Location: "Thao Dien, Ho Chi Minh City", PriceLabel: "$2.10M", Beds: 5, Baths: 4, Area: "420 sqm", Badge: "Villa", Image: "/assets/img/listing_3.avif", Summary: "A family villa with a tropical courtyard, private pool, and separate guest pavilion for extended stays.", Gallery: []string{"/assets/img/listing_3.avif", "/assets/img/listing_7.avif", "/assets/img/listing_11.avif"}, Headline: "A low-density garden villa with strong privacy and resort-style outdoor rooms.", Description: "This villa is arranged around a sheltered courtyard and pool edge, giving the home a quieter, more residential feel while keeping lifestyle amenities close by.", Options: []string{"Private pool and garden", "Guest pavilion", "Large family kitchen", "Multi-car frontage"}, Info: []DetailItem{{Label: "Property type", Value: "Villa"}, {Label: "Availability", Value: "Private treaty"}, {Label: "Parking", Value: "3 car bays"}, {Label: "Furnishing", Value: "Custom interiors"}}, Features: []string{"Large sliding openings to shaded outdoor dining", "Ground-floor guest suite for visiting family", "Primary bedroom overlooking pool landscape", "Excellent fit for long-term owner occupation"}},
		{Slug: "the-marq-executive-suite", Title: "The Marq Executive Suite", Category: "rent", Location: "District 1, Ho Chi Minh City", PriceLabel: "$3,600/mo", Beds: 2, Baths: 2, Area: "118 sqm", Badge: "Business", Image: "/assets/img/listing_5.avif", Summary: "Fully furnished executive suite steps from major offices, with concierge support and private wellness amenities.", Gallery: []string{"/assets/img/listing_5.avif", "/assets/img/listing_8.avif", "/assets/img/listing_12.avif"}, Headline: "A polished central-district rental for clients who value finish quality and convenience.", Description: "The Marq Executive Suite is positioned for tenants who want a fast commute, high service standards, and an interior that feels composed from day one.", Options: []string{"Move-in ready furnishing", "District 1 location", "Tenant concierge support", "Flexible lease structure"}, Info: []DetailItem{{Label: "Property type", Value: "Executive rental suite"}, {Label: "Availability", Value: "Immediate"}, {Label: "Lease term", Value: "12 months+"}, {Label: "Pets", Value: "By request"}}, Features: []string{"Hotel-style arrival lobby", "Quiet home office nook", "Wellness amenities within the building", "Strong option for relocation and corporate tenants"}},
		{Slug: "palm-marina-loft", Title: "Palm Marina Loft", Category: "rent", Location: "District 7, Ho Chi Minh City", PriceLabel: "$2,400/mo", Beds: 2, Baths: 2, Area: "132 sqm", Badge: "Loft", Image: "/assets/img/listing_6.avif", Summary: "A double-height loft with marina views, custom storage, and flexible work-from-home corners.", Gallery: []string{"/assets/img/listing_6.avif", "/assets/img/listing_8.avif", "/assets/img/listing_9.avif"}, Headline: "A more character-driven rental with height, light, and marina-facing atmosphere.", Description: "Palm Marina Loft uses vertical volume well, with layered storage, a strong living zone, and an adaptable upper level that supports hybrid work.", Options: []string{"Loft volume", "Marina outlook", "Custom storage", "Remote-work friendly"}, Info: []DetailItem{{Label: "Property type", Value: "Loft apartment"}, {Label: "Availability", Value: "Immediate"}, {Label: "Lease term", Value: "12 months+"}, {Label: "Furnishing", Value: "Fully furnished"}}, Features: []string{"Mezzanine layout for clearer zoning", "Natural light across the main living zone", "Waterfront promenades and cafes nearby", "Ideal for design-conscious tenants"}},
		{Slug: "canal-house-townhome", Title: "Canal House Townhome", Category: "sale", Location: "Binh Thanh, Ho Chi Minh City", PriceLabel: "$1.18M", Beds: 4, Baths: 3, Area: "246 sqm", Badge: "Townhome", Image: "/assets/img/listing_7.avif", Summary: "Light-filled townhome featuring layered garden terraces, a quiet study, and generous family gathering spaces.", Gallery: []string{"/assets/img/listing_7.avif", "/assets/img/listing_11.avif", "/assets/img/listing_10.avif"}, Headline: "A family townhome that layers private corners with social living rooms.", Description: "Canal House Townhome offers the flexibility of landed-style living in a more connected urban neighborhood, with terraces, family areas, and strong internal zoning.", Options: []string{"Terraced outdoor spaces", "Quiet home office", "Family-oriented planning", "Well-connected address"}, Info: []DetailItem{{Label: "Property type", Value: "Townhome"}, {Label: "Availability", Value: "By appointment"}, {Label: "Parking", Value: "2 car bays"}, {Label: "Furnishing", Value: "Partially furnished"}}, Features: []string{"Generous living room opening to planted terrace", "Study space separated from family activity", "Balanced scale for multi-generational living", "Good access to schools and central districts"}},
		{Slug: "the-opera-rental-collection", Title: "The Opera Rental Collection", Category: "rent", Location: "Ba Son, Ho Chi Minh City", PriceLabel: "$4,200/mo", Beds: 3, Baths: 2, Area: "168 sqm", Badge: "Prime", Image: "/assets/img/listing_8.avif", Summary: "Premium rental inventory for clients who want immediate move-in near dining, riverfront walks, and transit.", Gallery: []string{"/assets/img/listing_8.avif", "/assets/img/listing_5.avif", "/assets/img/listing_12.avif"}, Headline: "Prime district rental inventory for clients who want a central lifestyle and immediate access.", Description: "This collection is positioned for premium renters seeking polished interiors, strong amenity access, and a location that supports both work and leisure.", Options: []string{"Prime central location", "High-floor inventory", "Managed rental support", "Riverfront adjacency"}, Info: []DetailItem{{Label: "Property type", Value: "Premium apartment"}, {Label: "Availability", Value: "Selected units available"}, {Label: "Lease term", Value: "12 months+"}, {Label: "Furnishing", Value: "Turnkey"}}, Features: []string{"Close to dining and cultural destinations", "Concierge-backed resident experience", "Strong option for executive households", "High convenience with minimal setup time"}},
		{Slug: "jardin-compact-studio", Title: "Jardin Compact Studio", Category: "rent", Location: "Phu Nhuan, Ho Chi Minh City", PriceLabel: "$1,050/mo", Beds: 1, Baths: 1, Area: "54 sqm", Badge: "Starter", Image: "/assets/img/listing_9.avif", Summary: "Compact, polished studio with integrated millwork, ideal for first-time renters seeking a calm neighborhood.", Gallery: []string{"/assets/img/listing_9.avif", "/assets/img/listing_12.avif", "/assets/img/listing_6.avif"}, Headline: "A compact rental with a cleaner layout and calmer neighborhood rhythm.", Description: "Jardin Compact Studio is aimed at first-time renters and single professionals who want something composed, practical, and easy to maintain.", Options: []string{"Compact efficient plan", "Integrated millwork", "Lower monthly spend", "Neighborhood cafes nearby"}, Info: []DetailItem{{Label: "Property type", Value: "Studio apartment"}, {Label: "Availability", Value: "Immediate"}, {Label: "Lease term", Value: "6 to 12 months"}, {Label: "Furnishing", Value: "Fully furnished"}}, Features: []string{"Highly efficient storage planning", "Simple one-level layout", "Good fit for first-time independent renters", "Walkable daily convenience"}},
		{Slug: "sora-family-residence", Title: "Sora Family Residence", Category: "sale", Location: "Thu Duc, Ho Chi Minh City", PriceLabel: "$860K", Beds: 3, Baths: 2, Area: "149 sqm", Badge: "Family", Image: "/assets/img/listing_10.avif", Summary: "Three-bedroom residence with a child-friendly clubhouse, resort pool, and shaded arrival court.", Gallery: []string{"/assets/img/listing_10.avif", "/assets/img/listing_11.avif", "/assets/img/listing_2.avif"}, Headline: "A balanced family apartment in a community geared toward everyday ease.", Description: "Sora Family Residence is a practical owner-occupier option with clean planning, shared amenities, and a neighborhood that appeals to growing families.", Options: []string{"Family club amenities", "Resort-style pool", "Efficient 3-bedroom plan", "Future-ready district"}, Info: []DetailItem{{Label: "Property type", Value: "Family apartment"}, {Label: "Availability", Value: "Ready to move"}, {Label: "Parking", Value: "1 basement bay"}, {Label: "Furnishing", Value: "Semi-furnished"}}, Features: []string{"Bedrooms placed for privacy from the living room", "Good amenity support for younger households", "Shaded site planning and comfortable arrival", "Attractive for upgraders leaving central districts"}},
		{Slug: "the-lantern-garden-flat", Title: "The Lantern Garden Flat", Category: "sale", Location: "District 9, Ho Chi Minh City", PriceLabel: "$610K", Beds: 2, Baths: 2, Area: "104 sqm", Badge: "Garden", Image: "/assets/img/listing_11.avif", Summary: "A low-rise apartment overlooking a central garden spine, balancing greenery with quick commuter access.", Gallery: []string{"/assets/img/listing_11.avif", "/assets/img/listing_10.avif", "/assets/img/listing_7.avif"}, Headline: "A quieter low-rise apartment with greenery at the center of the living experience.", Description: "The Lantern Garden Flat offers a softer, more residential mood than higher-density towers, while still keeping useful connections to newer employment hubs.", Options: []string{"Low-rise setting", "Garden-facing outlook", "Efficient 2-bedroom planning", "Commuter-friendly district"}, Info: []DetailItem{{Label: "Property type", Value: "Garden apartment"}, {Label: "Availability", Value: "Ready to move"}, {Label: "Parking", Value: "1 basement bay"}, {Label: "Furnishing", Value: "Unfurnished"}}, Features: []string{"Gentler scale and lower-density feel", "Good balance between quiet living and access", "Suitable for end-users and smaller investors", "Strong daylight across the main room"}},
		{Slug: "harbour-edge-apartment", Title: "Harbour Edge Apartment", Category: "rent", Location: "District 4, Ho Chi Minh City", PriceLabel: "$1,900/mo", Beds: 2, Baths: 1, Area: "88 sqm", Badge: "Harbour", Image: "/assets/img/listing_12.avif", Summary: "River-adjacent rental apartment with efficient planning, layered lighting, and a breezy corner balcony.", Gallery: []string{"/assets/img/listing_12.avif", "/assets/img/listing_8.avif", "/assets/img/listing_9.avif"}, Headline: "A river-adjacent rental for tenants who want central access with a lighter monthly commitment.", Description: "Harbour Edge Apartment pairs a useful two-bedroom layout with a convenient location and a comfortable level of finish for long-term tenancy.", Options: []string{"Corner balcony", "River-adjacent location", "Efficient 2-bedroom plan", "Good value monthly rate"}, Info: []DetailItem{{Label: "Property type", Value: "Apartment"}, {Label: "Availability", Value: "Immediate"}, {Label: "Lease term", Value: "12 months+"}, {Label: "Furnishing", Value: "Partially furnished"}}, Features: []string{"Well-sized living zone for the segment", "Easy connection to District 1", "Good fit for couples or small households", "Practical base for city-focused renters"}},
	}
}

func Projects() []Project {
	return []Project{
		{Title: "Aster Bay Residences", Status: "launching", Location: "Thu Thiem, Thu Duc", Delivery: "Q2 2027", PriceLabel: "From $480K", Image: "/assets/img/project_1.avif", Summary: "A new waterfront address pairing wellness decks, work lounges, and high-floor river panoramas."},
		{Title: "Crescent Courtyard Homes", Status: "pre-launch", Location: "District 7, Ho Chi Minh City", Delivery: "Q4 2027", PriceLabel: "From $390K", Image: "/assets/img/project_2.avif", Summary: "Mid-rise family homes organized around shaded courtyards, community gardens, and flexible social spaces."},
		{Title: "Nova Heritage Tower", Status: "ready", Location: "District 1, Ho Chi Minh City", Delivery: "Ready to move", PriceLabel: "From $720K", Image: "/assets/img/project_3.avif", Summary: "Completed branded residences with concierge services, private dining rooms, and walkable city-center access."},
		{Title: "Lagoon District Villas", Status: "launching", Location: "Thu Duc, Ho Chi Minh City", Delivery: "Q1 2028", PriceLabel: "From $1.2M", Image: "/assets/img/project_4.avif", Summary: "Garden villas with lagoon-edge promenades, clubhouse amenities, and a quieter low-density neighborhood feel."},
		{Title: "The Vale Signature", Status: "ready", Location: "Binh Thanh, Ho Chi Minh City", Delivery: "Ready to move", PriceLabel: "From $530K", Image: "/assets/img/project_5.avif", Summary: "Contemporary ready inventory aimed at urban professionals wanting quick access to the central business core."},
		{Title: "Metro Leaf Quarter", Status: "pre-launch", Location: "District 9, Ho Chi Minh City", Delivery: "Q3 2028", PriceLabel: "From $310K", Image: "/assets/img/project_6.avif", Summary: "Transit-connected homes designed around park frontage, outdoor learning areas, and neighborhood retail."},
		{Title: "Canvas River Park", Status: "launching", Location: "District 2, Ho Chi Minh City", Delivery: "Q4 2027", PriceLabel: "From $610K", Image: "/assets/img/project_7.avif", Summary: "A mixed-use launch with stepped terraces, waterfront dining, and a broad resident amenity podium."},
		{Title: "Maison Verre", Status: "ready", Location: "Thao Dien, Ho Chi Minh City", Delivery: "Ready to move", PriceLabel: "From $880K", Image: "/assets/img/project_8.avif", Summary: "A limited collection of glass-wrapped residences designed for buyers who prioritize finish quality and privacy."},
		{Title: "Harborline Commons", Status: "pre-launch", Location: "District 4, Ho Chi Minh City", Delivery: "Q2 2028", PriceLabel: "From $295K", Image: "/assets/img/project_9.avif", Summary: "Value-led apartments with future retail activation, community sports areas, and accessible city connections."},
	}
}

func News() []Article {
	return []Article{
		{Title: "Why riverfront inventory is leading high-end demand this quarter", Category: "Market Update", Published: "April 2, 2026", Image: "/assets/img/project_10.avif", Summary: "Buyer demand is clustering around locations that blend wellness, views, and short commute patterns."},
		{Title: "How to compare launch-phase pricing against ready-to-move stock", Category: "Buying Guide", Published: "March 28, 2026", Image: "/assets/img/listing_13.avif", Summary: "A practical checklist for deciding when launch incentives outweigh the certainty of completed inventory."},
		{Title: "Design trends shaping family-oriented projects in Thu Duc", Category: "Project Watch", Published: "March 22, 2026", Image: "/assets/img/project_5.avif", Summary: "Developers are favoring layered outdoor spaces, learning zones, and flexible multi-generational layouts."},
		{Title: "Rental absorption stays strong in central districts", Category: "Rental Report", Published: "March 18, 2026", Image: "/assets/img/listing_14.avif", Summary: "Professionally managed apartments with amenities and quality interiors continue to outperform the broader market."},
	}
}

func FilterListings(query string, category string) []Listing {
	items := Listings()
	filtered := make([]Listing, 0, len(items))
	query = strings.TrimSpace(strings.ToLower(query))
	category = strings.TrimSpace(strings.ToLower(category))

	for _, item := range items {
		if category != "" && category != "all" && item.Category != category {
			continue
		}

		if query != "" {
			haystack := strings.ToLower(item.Title + " " + item.Location + " " + item.Badge + " " + item.Summary)
			if !strings.Contains(haystack, query) {
				continue
			}
		}

		filtered = append(filtered, item)
	}

	return filtered
}

func FilterProjects(query string, status string) []Project {
	items := Projects()
	filtered := make([]Project, 0, len(items))
	query = strings.TrimSpace(strings.ToLower(query))
	status = strings.TrimSpace(strings.ToLower(status))

	for _, item := range items {
		if status != "" && status != "all" && item.Status != status {
			continue
		}

		if query != "" {
			haystack := strings.ToLower(item.Title + " " + item.Location + " " + item.Status + " " + item.Summary)
			if !strings.Contains(haystack, query) {
				continue
			}
		}

		filtered = append(filtered, item)
	}

	return filtered
}

func FindListingBySlug(slug string) (Listing, bool) {
	for _, item := range Listings() {
		if item.Slug == slug {
			return item, true
		}
	}
	return Listing{}, false
}
