package tlds

import "sync"

var PopularDomains = []string{".com", ".co", ".net", ".org", ".info"}

func StreamPopularDomains() <-chan string {
	messages := make(chan string)
	go func() {
		defer close(messages)
		for _, d := range PopularDomains {
			messages <- d
		}
	}()
	return messages
}

var CountryDomains = []string{
	".ac", ".ad", ".ae", ".af", ".ag", ".ai", ".al", ".am", ".an", ".ao", ".aq", ".ar", ".as",
	".at", ".au", ".aw", ".ax", ".az", ".ba", ".bb", ".bd", ".be", ".bf", ".bg", ".bh", ".bi",
	".bj", ".bm", ".bn", ".bo", ".bq", ".br", ".bs", ".bt", ".bv", ".bw", ".by", ".bz", ".ca",
	".cc", ".cd", ".cf", ".cg", ".ch", ".ci", ".ck", ".cl", ".cm", ".cn", ".co", ".cr", ".cu",
	".cv", ".cw", ".cx", ".cy", ".cz", ".de", ".dj", ".dk", ".dm", ".do", ".dz", ".ec", ".ee",
	".eg", ".eh", ".er", ".es", ".et", ".eu", ".fi", ".fj", ".fk", ".fm", ".fo", ".fr", ".ga",
	".gb", ".gd", ".ge", ".gf", ".gg", ".gh", ".gi", ".gl", ".gm", ".gn", ".gp", ".gq", ".gr",
	".gs", ".gt", ".gu", ".gw", ".gy", ".hk", ".hm", ".hn", ".hr", ".ht", ".hu", ".id", ".ie",
	".il", ".im", ".in", ".io", ".iq", ".ir", ".is", ".it", ".je", ".jm", ".jo", ".jp", ".ke",
	".kg", ".kh", ".ki", ".km", ".kn", ".kp", ".kr", ".kw", ".ky", ".kz", ".la", ".lb", ".lc",
	".li", ".lk", ".lr", ".ls", ".lt", ".lu", ".lv", ".ly", ".ma", ".mc", ".md", ".me", ".mg",
	".mh", ".mk", ".ml", ".mm", ".mn", ".mo", ".mp", ".mq", ".mr", ".ms", ".mt", ".mu", ".mv",
	".mw", ".mx", ".my", ".mz", ".na", ".nc", ".ne", ".nf", ".ng", ".ni", ".nl", ".no", ".np",
	".nr", ".nu", ".nz", ".om", ".pa", ".pe", ".pf", ".pg", ".ph", ".pk", ".pl", ".pm", ".pn",
	".pr", ".ps", ".pt", ".pw", ".py", ".qa", ".re", ".ro", ".rs", ".ru", ".rw", ".sa", ".sb",
	".sc", ".sd", ".se", ".sg", ".sh", ".si", ".sj", ".sk", ".sl", ".sm", ".sn", ".so", ".sr",
	".ss", ".st", ".su", ".sv", ".sx", ".sy", ".sz", ".tc", ".td", ".tf", ".tg", ".th", ".tj",
	".tk", ".tl", ".tm", ".tn", ".to", ".tp", ".tr", ".tt", ".tv", ".tw", ".tz", ".ua", ".ug",
	".uk", ".us", ".uy", ".uz", ".va", ".vc", ".ve", ".vg", ".vi", ".vn", ".vu", ".wf", ".ws",
	".ye", ".yt", ".za", ".zm", ".zw"}

func StreamCountryDomains() <-chan string {
	messages := make(chan string)
	go func() {
		defer close(messages)
		for _, d := range CountryDomains {
			messages <- d
		}
	}()
	return messages
}

var GenericDomains = []string{
	".academy", ".accountant", ".accountants", ".active", ".actor", ".adult", ".aero",
	".agency", ".airforce", ".apartments", ".app", ".archi", ".army", ".associates",
	".attorney", ".auction", ".audio", ".autos", ".band", ".bar", ".bargains", ".beer",
	".best", ".bid", ".bike", ".bingo", ".bio", ".biz", ".black", ".blackfriday", ".blog",
	".blue", ".boo", ".boutique", ".build", ".builders", ".business", ".buzz", ".cab", ".cam",
	".camera", ".camp", ".cancerresearch", ".capital", ".cards", ".care", ".career", ".careers",
	".cars", ".cash", ".casino", ".catering", ".center", ".ceo", ".channel", ".chat", ".cheap",
	".christmas", ".church", ".city", ".claims", ".cleaning", ".click", ".clinic", ".clothing",
	".cloud", ".club", ".coach", ".codes", ".coffee", ".college", ".community", ".company",
	".computer", ".condos", ".construction", ".consulting", ".contractors", ".cooking", ".cool",
	".coop", ".country", ".coupons", ".credit", ".creditcard", ".cricket", ".cruises", ".dad",
	".dance", ".date", ".dating", ".day", ".deals", ".degree", ".delivery", ".democrat", ".dental",
	".dentist", ".design", ".diamonds", ".diet", ".digital", ".direct", ".directory", ".discount",
	".dog", ".domains", ".download", ".eat", ".education", ".email", ".energy", ".engineer",
	".engineering", ".equipment", ".esq", ".estate", ".events", ".exchange", ".expert", ".exposed",
	".express", ".fail", ".faith", ".family", ".fans", ".farm", ".fashion", ".feedback", ".finance",
	".financial", ".fish", ".fishing", ".fit", ".fitness", ".flights", ".florist", ".flowers", ".fly",
	".foo", ".football", ".forsale", ".foundation", ".fund", ".furniture", ".fyi", ".gallery", ".garden",
	".gift", ".gifts", ".gives", ".glass", ".global", ".gold", ".golf", ".gop", ".graphics", ".green",
	".gripe", ".guide", ".guitars", ".guru", ".healthcare", ".help", ".here", ".hiphop", ".hiv", ".hockey",
	".holdings", ".holiday", ".homes", ".horse", ".host", ".hosting", ".house", ".how", ".info", ".ing",
	".ink", ".institute", ".insure", ".international", ".investments", ".jewelry", ".jobs", ".kim", ".kitchen",
	".land", ".lawyer", ".lease", ".legal", ".lgbt", ".life", ".lighting", ".limited", ".limo", ".link",
	".loan", ".loans", ".lol", ".lotto", ".love", ".luxe", ".luxury", ".management", ".market", ".marketing",
	".markets", ".mba", ".media", ".meet", ".meme", ".memorial", ".men", ".menu", ".mobi", ".moe", ".money",
	".mortgage", ".motorcycles", ".mov", ".movie", ".museum", ".name", ".navy", ".network", ".new", ".news",
	".ngo", ".ninja", ".one", ".ong", ".onl", ".online", ".ooo", ".organic", ".partners", ".parts", ".party",
	".pharmacy", ".photo", ".photography", ".photos", ".physio", ".pics", ".pictures", ".pid", ".pink", ".pizza",
	".place", ".plumbing", ".plus", ".poker", ".porn", ".post", ".press", ".pro", ".productions", ".prof",
	".properties", ".property", ".qpon", ".racing", ".recipes", ".red", ".rehab", ".ren", ".rent", ".rentals",
	".repair", ".report", ".republican", ".rest", ".review", ".reviews", ".rich", ".rip", ".rocks", ".rodeo",
	".rsvp", ".run", ".sale", ".school", ".science", ".services", ".sex", ".sexy", ".shoes", ".show", ".singles",
	".site", ".soccer", ".social", ".software", ".solar", ".solutions", ".space", ".studio", ".style", ".sucks",
	".supplies", ".supply", ".support", ".surf", ".surgery", ".systems", ".tattoo", ".tax", ".taxi", ".team",
	".store", ".tech", ".technology", ".tel", ".tennis", ".theater", ".tips", ".tires", ".today", ".tools", ".top",
	".tours", ".town", ".toys", ".trade", ".training", ".travel", ".university", ".vacations", ".vet", ".video",
	".villas", ".vision", ".vodka", ".vote", ".voting", ".voyage", ".wang", ".watch", ".webcam", ".website", ".wed",
	".wedding", ".whoswho", ".wiki", ".win", ".wine", ".work", ".works", ".world", ".wtf", ".xxx", ".xyz", ".yoga", ".zone"}

func StreamGenericDomains() <-chan string {
	messages := make(chan string)
	go func() {
		defer close(messages)
		for _, d := range GenericDomains {
			messages <- d
		}
	}()
	return messages
}

func StreamAllDomains() <-chan string {
	messages := make(chan string)
	go func() {
		defer close(messages)
		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			defer wg.Done()
			for d := range StreamCountryDomains() {
				messages <- d
			}
		}()
		wg.Add(1)
		go func() {
			defer wg.Done()
			for d := range StreamGenericDomains() {
				messages <- d
			}
		}()
		wg.Add(1)
		go func() {
			defer wg.Done()
			for d := range StreamPopularDomains() {
				messages <- d
			}
		}()
		wg.Wait()
	}()
	return messages
}
