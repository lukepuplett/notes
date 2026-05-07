# LinkedIn

---

##### linkedin post todo



# Rethinking AI Tools: Why We're Reinventing the Wheel and How the Web Already Solved It



Hey everyone, I've been diving deep into how AI agents interact with tools and APIs lately, and it's got me thinking about something that's been bugging me for a while. You know those protocols like MCP (Model Context Protocol) that Anthropic put out? They're basically trying to standardize how LLMs call functions, fetch data, and plug into external systems. On the surface, it sounds great, a kind of universal connector for AI. But the more I look at it, the more it feels like we're just recreating stuff the web figured out decades ago, and not in the best way.



Let me back up a bit. The web, at its core, is all about content paired with actions. A webpage isn't just text or images; it's a context with built-in tools like links and forms. When you're on a site editing your profile, you only see the relevant fields and buttons for that task. Click a link, and you're navigated to a new context with its own set of options. Forms are genius here: they tell you where to send data (the action URL), what parameters to include (input fields), and even possible values (like dropdowns). It's all self-contained, discoverable, and keeps things focused and efficient.



Now, contrast that with MCP. It's essentially JSON-RPC over various transports, where tools are discovered upfront through schemas, and the LLM decides what to call based on its reasoning. Sure, it handles things like runtime plugging in new tools or streaming results, which is handy for agents. But it reinvents links and forms in a clunky way. There's no native equivalent to a redirect shifting context seamlessly, or cookies handling auth without extra hassle. Hypermedia APIs have been doing this in JSON for years, but even they struggle because JSON doesn't have a standard for forms. MCP feels like a step sideways, born from trying to make LLMs reliable at function calling without leaning on the web's proven patterns.



I suspect a big reason for this is skill silos in development teams. From my experience as a backend engineer who's also spent time in web and hypermedia land, backend folks love structured schemas and explicit RPC calls. It's clean, typed, and maps to what we build every day. But if you're steeped in REST as true hypermedia (HATEOAS style, where the API drives state transitions through links and controls), you see the web as the ultimate model for discoverable interfaces. MCP's design screams "backend engineers solving AI pain points" more than "web natives optimizing for evolvability." If more web-savvy people had been in the design room, we might have ended up with something that treats HTML as the tool description language, with LLMs parsing pages and outputting HTTP requests directly.



Imagine that alternative: an LLM attached to your browser or chat app that reads the current page's content, spots the forms and links, and acts on them. You ask it to do something, it matches your intent to an affordance, crafts a POST or GET request, and the client executes it. Responses come back, redirects switch contexts, cookies handle sessions and login flows. It's stateless where it needs to be, with URLs encoding journey details reliably. No need for a whole new protocol; just train models to grok HTML better and let a thin client handle the transport. This would leverage existing hypermedia APIs, make auth straightforward (find the login form, submit creds, store the cookie), and feel way more natural.



And here's where it gets exciting: this setup could revolutionize how we browse and interact with apps. An AI trained on HTML like this wouldn't just fetch and dump everything; it could intelligently select the most relevant parts of a page and display only those in its own compact chat window or mini-browser view. Think about it: no more wading through ads, sidebars, or irrelevant nav. You query for flight options, and it extracts just the results table, filters, and booking buttons, rendering a clean snippet with preserved interactivity. It could annotate why it picked those bits, collapse sections, or reformat for better readability. Browsing shifts from passive scrolling to goal-directed extraction, with the AI as your personal UI curator.



Of course, this isn't without hurdles. One big one is handling dynamic JS-heavy sites, where static HTML doesn't show the full picture; you'd need better rendering or accessibility parsing. Security is another: sandboxing to prevent malicious submissions. But the real showstopper for many apps is legally important content. Disclaimers, terms of use, privacy notices, or regulatory warnings can't be filtered as "noise." Stripping them could lead to liability issues, misinformation, or compliance violations like GDPR.



To make this future work, we'd need a thoughtful system. Start with standardized markup: new HTML elements or attributes like role="legal-disclaimer" or meta tags flagging essential sections with CSS selectors. Site owners could declare what must be preserved, building on things like schema.org or robots.txt. Then, AI smarts: train models to recognize legal patterns (all-caps blocks, specific phrases, common placements) and always include them in renders, maybe as persistent footers or expandables. Client-side enforcement helps too: the chat window applies policies to fetch and show declared legals, with user toggles for full views or summaries. Finally, regulatory support: guidelines from bodies like the W3C or FTC to mandate AI-safe rendering, with safe harbors for compliant systems.



There are trade-offs, like over-preservation diluting curation or adversarial sites abusing tags. Global laws vary, so location-aware logic is key. But with consensus on standards, this could scale safely. Sites might add these markers to optimize for AI traffic, just like SEO today.



In the end, MCP isn't dumb; it solves real integration pains in the agent world. But by leaning harder on the web's foundations, we could build simpler, more composable AI experiences. Models are getting better at HTML every day, so maybe the "LLM as browser brain" era is closer than we think. What do you folks reckon? Is this the path forward, or are dedicated protocols like MCP here to stay? Drop your thoughts below.



##### LinkedIn Post



Luke Puplett

@lukepuplett



Aaron and @GergelyOrosz are describing the last 35 years but on a condensed timescale.



Even on the slow, organic track of the last decades, most companies struggled to make the transition to software. We instead had new companies form and take over from California.



It needed new companies because transitioning to a design and engineering company meant rescinding the twentieth century corporate stack; skyscrapers with Wall St. friendly accountants at the top and suits holding power. Existing companies are unable to make this change, since the leaders are part of the suited class and will never abdicate and cede control to the hoodies.



Fun note: Apple's HQ is a ring because it thus cannot have corner offices.



European coders experienced this resistance as the failure of The Agile Manifesto where agile was mutilated and perverted into an SOP for corporate micromanagement.



It's akin to @nfergus Square and the Tower. California was the only place where new, flat companies of network effects could be bootstrapped and thrive, all downstream of Arthur Rock and Intel.



Now, with AI, we have Silicon Valley on nitrous. But an interesting dynamic is hinted at by 

@levie

 



"The limiter then becomes how rapidly your customers can actually adopt new software"



This was always the limiter! Software engineers have always been able to see 10 or 20 years into the future, using knowledge of their tools and customer problems, and envision and build dream solutions, but they cannot get people to use them, because tech outsiders, i.e. prospective customers, are only human and don't know what they don't know.



Not to mention that most buyers of software are scared of getting it wrong and getting fired, because most companies don't tolerate mistakes and are highly gamified environments. Brilliant internal solutions can give techies too much power. And finally, customers can be scared of job losses from automation and diminishment of their empires.



So customers have huge status quo bias and prefer faster horses. This is the timing problem, or the 'build something people want' problem, or 'crossing the chasm'. Being correct at the wrong time is the same as being wrong. You can only hit the market with your crazy idea when the 'ordinary' masses no longer think it's all that crazy.



Now, the other dynamic at play this time around is just how much program code is going to be needed at all.



Once agents are doing the work of coded workflows and business logic, and agents are doing the work of interfacing with users, and doing the work of pulling insights from files and databases, writing and sending documents, filing compliance reports, and knowing all the things and having all the skills, plus, showing any remaining humans their reports, tables, charts and decision supporting materials, then aren't we all racing with AI written code to pick up pennies in front of the AI driven steam roller?



That is the elephant in the room.



The race then is to fill the gap between now and the crazy future with hybrid traditional software before customers get too comfortable and want that future.



