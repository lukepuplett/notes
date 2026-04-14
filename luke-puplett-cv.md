# Luke Puplett - Versatile & Experienced Engineer

Hello - I'm a curious, hard-working and collaborative engineer with over two decades experience in very large and very small companies. You'll see a few themes in the stories below; ingenuity, agency and improving how teams build stuff.

I've invented and delivered award-winning tools for managing IT at big banks, built visually rich multi-screen trading apps, tried to build a predictive analytics startup, developed a new app for medics in the US, ran a Python supercomputer and wrote an API auth system that my client patented. I also designed and built my own house, dropped out of one of England's best secondary schools when I was only 13 to write code, and my first job was in Microsoft's tech support team.

Coming from an ops background in vast banking environments where I invented tools for automation and control - SRE before it had a name - I've broad knowledge and interests spanning design and UX to distributed architectures, low-latency trading and lock-free, low-GC code, but also very much company culture and organising people and work.

I work in small batches. I minimise time-to-insight. I automate all the things!! I slow down to speed up. I read voraciously and make time to learn. I teach those around me. I write for others. I focus. I research how best to do something, then I strap in and crack on with it.

This CV is the super long director's cut. If you have a similarly long specification of your role and ideal person, I recommend pasting my CV and your spec into an LLM and asking it if I might be a good or bad fit and why - I think it's the future.

## Experience

**January 2020 - Present Day**

## Founder and Engineer at Zipwire

*Recruiter's Paradise - Automating Payments, Identity, Blockchain Finance*

#### Context and Business Situation

Zipwire Approve is a worker’s private journal, a WhatsApp bot and a Go CLI for techies used to fill timesheets according to rules so they cannot be wrong, with workflows for approvals and payouts. Zipwire Collect gathers and scans documents and evidence over the web or WhatsApp, runs ID checks using gov-approved selfie tech. Zipwire Attest can witness these facts on Ethereum.

This all feeds into the idea of putting invoice factoring onchain and automating agency backoffice cashflow and payouts so engineers get paid in minutes from their CLI.

ProofPack is an OSS NPM package, a JS and .NET SDK, I created for making and checking cryptographic selective-disclosure proofs. It’s a format I invented based around a JWT. Banks could use it to e.g. let customers download proofs of address or outgoings.

I started Zipwire during the pandemic while my wife became a headteacher and I looked after our young kids. I discovered the factoring headache and saw how my work on real-time insurance markets combines with blockchains to neatly solve this.

The timing is off; slow regulation, negative reputation of crypto, fragmentation and complexity. My solution hinges on several shifts that have not aligned yet.

> "The market may stay irrational longer than you can stay solvent."

#### Engineering Achievements

I learned Google Cloud, React (and SvelteKit), Docker, AI/LLMs, tool/function calling, MCP, even Hardhat and Rust but settled on C# and Golang. It’s a ‘semi-monolithic’ ASP.NET 6 MPA using Bootstrap, Sass and Knockout.js + JSON APIs and some HTMX. I began on Azure using blobs and Cosmos, DevOps/Pipelines before moving it to GCP on Docker in Cloud Run, built on Cloud Build and uses GCS, Firestore and BigQuery.

Zipwire is fully-featured, way too much to cover. It consumes several LLMs, Google Document AI, OpenAI, Stripe, Twilio, Yoti, FreeAgent, EAS and so on. Everything was built, designed or written by me alone. It uses a modular, eventually consistent ‘microservices in a monolith’ architecture with an in-process bus.

The aim is to build a DeFi auction house for invoices so I’ve been studying ‘mechanism design’, game theory, Myerson and Milgrom, lots of financial regs, I even rewrote half the FCA Handbook. In March 2025 alone I cranked out a 35,000 line Ethereum library with only AI autocomplete. It’s on [GitHub](https://github.com/luke-puplett/Evoq.Ethereum)/Nuget.org as [Evoq.Ethereum](https://github.com/luke-puplett/Evoq.Ethereum). Hand built recursive binary encoders and using Elliptic Curve libs.

#### Other

I also hacked together StyleShift using Go, JS and AI models, a Chrome extension which lets people restyle images on property websites - now a built-in feature on RightMove (UK). Some contract work for a small business in my local town, app integration work.

**June 2019 - October 2019**

## Cloud Architect at Incited for Various Insurers

*InsureTech - Big Data Applications and Azure Consultancy*

#### Context and Business Situation

Incited is a startup of experienced Azure devs and data scientists solutioning for UK insurers. We worked closely with clients on cloud-native data pipelines, stream processing for fraud and realtime analytics.

One client ingested up to 50mln big JSON quotes a day from comparison portals and our systems looked for patterns in pricing, renewals, churn and fraud. We'd build on Azure PaaS and hand it over to the client with training and perhaps help hiring data talent.

#### Engineering Achievements

I redesigned the on-prem, SOAP-based, real-time B2B marketplace that connects many brokers and insurers, 'Polaris imarket' for a cloud native, HA solution.

All Azure PaaS I proposed active-active regions, serverless bits, plus a redesigned stateless JSON API to be hypermedia driven with webhooks. Designed for durability of message data in the event of a regional outage as well as end-to-end security.

I began testing components using Azure Event Hubs, Stream Analytics, Service Bus Queues and Functions measuring their scaling capacity and failure modes. I submitted feedback to the Microsoft docs, engaged their PMs and even jumped on the odd call to Redmond.

**April 2018 - May 2019**

## Senior Systems Analyst and Architect at Equiniti for Canada Life

*Greenfield Pensions and Investments e-Commerce Product Team*

*Link Please read Retirement Advantage below before reading this.*

#### Context and Business Situation

In 2018 the lead on the Retirement Account team asked me to rejoin the team I’d worked with in 2015.
The underwriting business had been sold to Canada Life and Equiniti had bought IT/ops, devs and product teams, sadly severing product from business. The 70-strong team were doing fake-agile ‘waterscrumfall’ with a big relaunch months out. I came in as the main architect within System Analysis, mentoring and providing tech leadership.
The dev Tech Lead and I evangelised trunk-based dev; continuous tiny deployments to a set of live services with regular user feedback and instrumentation to guide us. We lost that pitch. The original microservices app was rapidly built on a "bet the company" basis and Canada Life wanted big changes.
Since the main Angular quoting app and its backend weren’t designed to be evolved, we rebuilt it around new services using CQRS and event sourcing.
The Tech Lead had run with the things I started in 2016 and had put together a great team of 16 devs. We moved to Azure Service Fabric (like Kubernetes) and integrated the Ops team into Dev. I challenged inertia and silos, tried to map impacts and align Equiniti’s incentives with Canada Life’s outcomes.

#### Engineering Achievements - [See detailed technicals](/me/equinity-technicals.md)

With a distant launch date and little hope of winning hearts around production pushes, the Tech Lead and I instead solutioned a total replacement of the existing system with many small event-driven services using RabbitMQ and a new Angular front-end.

We built eventually-consistent feature stacks - a really sweet design.

I introduced backends-for-frontends to decouple microservices from UI, insulating against front/backend flux and aiding the transition to event choreography.

I mentored 8 analysts and up to 16 developers on designing a proper, self-documenting HATEOAS application using JSON APIs the likes of which I have never seen any company produce. I worked on a PoC for a hypermedia client that I’ve since rewritten and open-sourced as [Surfdude](https://github.com/luke-puplett/Surfdude) - it self-drives hypermedia APIs.

And I designed a SignalR service to push events to the browser. Introduced Application Insights for observability and designed a tracing strategy. Tweaked Auth middleware.

Moved the team to a wiki and ran ‘Learn Markdown’ workshops for non-techies so we could centralise stories, designs and link it all.

Built a PoC for BDD to specify APIs using SpecFlow/Cucumber which drove the app via the hypermedia API via my Surfdude client - seriously cool, pretty sure a world first.

"You've been a massive influence for me professionally and I suspect you're unaware of the number of doors you've opened, and light bulb moments you've provided. This was both in the 2 weeks we initially worked at RA and on a more ongoing basis since you've come back into EQ. I'd like to say thank you and that I'm really, really grateful for all the ways you've helped me." - Technical Lead

**February 2017 - February 2018**

## Technical Lead at Centrica Energy Markets & Trading

*Scaling Python, Energy Options Trading, Valuation & Risk HPC Platform*

#### Context and Business Situation

Within the owner of British Gas sits Centrica EM&T, its prestigious unit that buys, sells and hedges a vast portfolio of energy assets. I ran a small team of devs that worked closely with Quantitative Analytics to value and stress test their nonlinear derivatives portfolios using Python quant models.

Our combined systems collated data from LNG cargoes, wind and gas turbines, FX and many other market feeds, then broke the problem into a DAG to run on thousands of compute cores on Azure overnight. We built and operated the harness, scheduling, data feeds and prep, business facing apps and monitoring, while the Quants built the models using NumPy, C++ and so on.

Traders accessed our HPC grid via Excel add-in or a web portal to which they submitted and monitored jobs. The on-prem side comprised several Windows services and an Angular SPA backed by an JSON/HTTP API. The Azure PaaS side was highly-available, active-active across two environments providing four deployments over two Azure regions so jobs could be moved between. A huge consumer of Azure Batch, I had the personal mobile of the PM in Seattle since we often hit regressions in Azure before anyone else.

We were often scratching our heads at CSV files at 3am, so I rented a flat near the office. The British Gas business had imposed heavy change control on essentially a fast moving quant hedge fund, so it was hard to iterate on reliability. The quants, whose models would crash on corrupt data, were not on pager duty so they had less incentive to write code that didn’t wake them up! If you build it, you run it.

#### Engineering Achievements - [See detailed technicals](/me/centrica-technicals.md)

I oversaw the development and launch of a trade version control system that enabled PnL changes to be "explained", like Git Blame for but derivatives trades - not my idea, but it was clever. Then I began breaking their monoliths and working with a Microsoft MVP on all the CICD automation using VSTS and Azure DevTest Labs, ARM templates, etc.

Introduced Scrum and a huge Kanban board. EM&T was breaking away from British Gas, moving to Azure, being ‘more Agile’. It took four weeks to get my PC setup with Visual Studio and 9 months to get some whiteboards. Code quality and engineering was a bit too rapid and the firefighting ate into the time needed to engineer a way out.

I designed much of their MiFID2 reporting solution using Azure Functions, Web Apps, Queues and SQL Data Warehouse. And for a while I worked on a greenfield .NET Core 2.0 proper REST API to maintain portfolio stress testing/what-if analysis on Azure App Service and CosmosDB.

Trained at Splunk’s London HQ for an observability solution for the distributed HPC environment. Splunk, like any high-perf database, needs custom-spec co-lo servers, but outsourced server provisioning can only put out cookie-cutter VMs.

**On Reflection**

The culture was a high-pressure and fairly political tradefloor where it could be difficult to apply yourself thickly on reliable engineering, a common footgun in trading.

**April 2015 - October 2016**

## Senior Developer and Security Expert at Retirement Advantage

*Greenfield Pensions and Investments e-Commerce Product Team*

#### Context and Business Situation

I was initially hired simply to write a SAML 2.0 SSO solution for a flexible new pension product to be used by tens of thousands of public users, but they were in dire need of delivery expertise having never built software like this. They'd been cargo cult software engineering for months, producing Gantt charts and piles of Word docs and little code. No process. No source control. No UX. No stories. They had 6 manual testers but no buildable app; they needed help but didn't even know it.

#### Engineering Achievements - [See detailed technicals](/me/mgm-tra-technicals.md)

With the help of a couple of other good contractors we delivered a slick single-page app. Despite having no authority, I steered them towards Git, Visual Studio Online, Slack, a semblance of Scrumban, NuGet, TeamCity, CI and towards the end, PowerShell DSC and even infra-as-code!

The big bang release was a crazy approach in this era so I bought the CEO The Lean Startup. I bought The Phoenix Project for the head of IT and their Head of Solutions Architecture. A Scrum book for the project manager. I tried to convince them of the power of working in small batches and decoupling deploy from release and a continually-deploying set of tested microservices. I tried so hard to explain that the APIs are the app and the BAs should be spec'ing the APIs. I was the only guy that really knew REST, microservices, CICD. It was tough and some developers came in and just left again.

We used Angular on the front-end written mostly by a creative agency, Web API 2.0 on the back with 5 primary microservices. I was officially responsible for security and user-management parts, including encryption, registrations, SSO, AD authentication, IIS and Windows hardening. I built a custom token and API permission and roles system which they patented, and set up a NuGet server. I was eventually allowed to set up a CICD pipeline (!) which was a revelation to all of the product and tech team, though they still insisted on humans manually copying files to production.

MSSQL with EF, MongoDB, and the app interfaced with an off-shelf pensions BO product via SOAP. The most interesting technical take-away was how APIs can be designed purely around GET and POST, and the natural emergence of a pattern we now call Backends for Frontends.

I was kept on for an extra month when their Solvency II team failed a regulatory deadline, so I automated and CI'ed complex SSIS/MSSQL using PowerShell+TeamCity.

#### On Reflection

As of end 2023 the app had been rebuilt three times but the 2016 auth service I wrote is still going strong - it's a cool design and could be a business on its own like Okta. My proudest achievements were the patent but also that two pen-testing companies found nothing, and that by the end, the SAs and BAs had gone from using Word to hacking on Postman :).

**December 2014 - February 2015**

## Developer and Architect at SunGard for Serco

*Data Analysis and Tooling for Large-scale IT*

#### Context and Business Situation

Hired by recommendation based on my highly innovative work in IT at Société Générale, see below. The government project was canned.

**October 2012 - December 2014**

## Consultant .NET Developer at OverStory for Springer Publishing

*Lean Startup & Horizon 3 Experiments in Healthcare and Media*

#### Context and Business Situation

OverStory is a big data and XML consultancy focusing on MarkLogic XML Server. I worked in a global, remote team on the Springer account led from the New York office. I was the lead C# developer in a small experimental Horizon 3 innovations team. We built new B2C and B2B e-commerce products on .NET, Java, XSLT and XQuery with HTTP APIs backed by a many-terabyte MarkLogic HA cluster containing all Springer's books, journals and media.

#### Engineering Achievements

Consulting on the development of a new healthcare product for front-line medical practitioners, trialling in US hospitals, our product took the form of a responsive, mobile-first website. Used HTML5, CSS3 and JavaScript (sparingly) on the Zurb Foundation framework for a fast in/out experience in a very low-bandwidth setting. The MVP validated key assumptions and the product won further investment and resources.

Working with the e-Product Manager, 3 BAs, 2 testers, up to 3 junior programmers, 1 ops/support engineer and another OverStory consultant in Canada, I was solely responsible for the web and .NET implementation providing thought-leadership on technology and architecture. Built a custom CMS on ASP.NET MVC 4 and specified the XSD schema, tooling and docs for an offshore team to translate and upload, rich media, medical documents.

I also provided the initial design direction and wire-framing, working with the primary stakeholder in Switzerland and her VP in New York, helping with keeping the product Lean. Helped with the full lifecycle, from inception to user journeys, persona creation and turning it all into backlog items and Kanban cards.

Overhauled their problematic legacy sites adding instrumentation and diagnostics and reducing CPU utilization from 90pct across the farm to under 4pct. Setup Git, TeamCity and NuGet, and coded a full "push-button" deployment pipeline via PowerShell to Windows and Linux. Built new APIs for new tablet apps on WebApi 2.

Automated everything. Setup a private NuGet feed for internal dependency management. Got 10 websites and services fully DevOps-ified and instantly deployable/rollbackable. We presented an SOA and microservices architecture for all future projects. Use of NLBs, CDNs, AWS and Mashery to provide security and availability. We developed REST microservices in an agile way between teams and geographies using stubs and a neat online tool called Apiary.

**On Reflection**

This was one of the most inspiring contracts I ever had. The old analogue ship was turned around by a bright New Yorker, Brian Bishop, who had brought in ThoughtWorks to help transition to agile, real agile. The level of autonomy was amazing, work anywhere culture, flexible working hours, the ability to hire expertise anywhere in the world, the use of whatever software on the web we wanted to organise ourselves, whatever cloud provider was the best, the most suitable databases and languages.

I got the break-up and automation to a point where we pushed many little changes live as soon as a QA had moved the card left. It was completely liberating and I've been fascinated by their open, safe, enabling culture ever since.

Note that from this point on I say less about the technologies used unless they're still relevant or particularly interesting.

**May 2012 - October 2012**

## Senior .NET Developer at Société Générale

*Large-scale Operations Engineering in Investment Banking*

#### Context and Business Situation

These days this team might be called Site Reliability Engineering; (SRE) is a discipline that incorporates aspects of software engineering and applies them to IT operations problems.

I returned to SocGen to work within the EMEA Operations and Engineering team on a small greenfield project of my own design to house data about their 40,000 PC desktop estate.

#### Engineering Achievements

Designed and built a web API that abstracted and represented WMI and AD. This allowed other tools built for non-technical and lesser-authorised staff to interact with the entire Windows estate and Active Directory and alleviate workload on high-cost skill teams. For example, you could replace a registry key like so PUT { // JSON } > https://my.api/{hostname}/HKLM/Software/Microsoft/...

Built many PowerShell CmdLets that consumed my APIs so the ops people could get stuff done productively. I also updated a much-loved Excel add-in I wrote in 2006 to use managed code and my new web APIs.

**On Reflection**

The things I built were cool, especially to the engineers who got the benefit of using them, but the most interesting thing about working alone in a team like this is the sustained productivity. I deployed all day long, getting immediate feedback from those around me and added an average of 441 lines of code a day. Every operations team should have a developer/inventor on board.

**November 2011 - May 2012**

## Entrepreneur, Architect & Developer for Myself

*Social Media & Data Analytics Startup*

*Link Please read about vuPlan.tv™ below before reading this.*

#### Context and Business Situation

Between freelance work I built and ran vuPlan.tv, a mobile and a desktop website with branded mobile apps and desktop clients.

#### Engineering Achievements

Added Facebook login to the phone app with hand-coded OAuth 2.0 and FB Graph API interaction. Published two apps to the Windows Phone marketplace and went live to the UK & Ireland in February 2012. My WPF client for Windows was live and gathering data.

**On Reflection**

I did all the design, architecture, branding, artwork, coded over 160,000 lines, hardware builds, hosting, security, PR, pitch decks, the privacy policy – everything myself!

I tried and I failed. There are too many lessons to discuss here but ultimately I didn't know what Lean Startup was so I failed very late.

**July 2011 - November 2011**

## Senior WPF Developer at Alpha Kinetic

*Derivatives Trading FinTech Startup*

#### Context and Business Situation

This start-up had a trade capture and risk product called Glide built in WPF that they sold to boutique hedge funds. They saw that I'd built Arena (below) and eagerly hired me after a funding round to build the product out and introduce some rigour and professionalism.

The CTO almost immediately walked out in a row over his stock. The engagement fizzled and the company was sold. It's not interesting enough to warrant your time.

**January 2011 - July 2011**

## Senior WPF Developer at Baker Technology for Lloyds Bank Corporate Markets

*Large FX Derivatives Trading Platform*

#### Context and Business Situation

Morgan Stanley's 'Matrix' single dealer platform built in Flex/Flash precipitated a race to build the most visually awesome web dealing experience. I was a key member of an 8-person front-end team delivering Lloyds Arena, a very high-profile answer to Matrix in just 6-months.

#### Engineering Achievements

I designed the majority of the application extensibility/plug-in framework. The application was composed of individual, loosely-coupled UI components and services, loaded dynamically using MEF. Each piece communicated via a simple in-process message bus.

The event-based coupling between components within the monolithic, GUI codebase is the same principle as for distributed microservices, with the same benefits! The team worked independently on their own components with little team interaction needed, only events to consume as necessary.

I was solely responsible for building the custom drag-drop window-docking control which was a key product differentiator and meant we had no dependencies on 3rd party controls. Troubleshooting memory leaks and reducing GC pauses was a significant challenge, as too was coding for multithreading and re-entrancy before the Task class had been "invented", and long before async-await. I built a rudimentary Task-like object based on an example by Joe Duffy.

**On Reflection**

My first £multi-million grown-up development role, I learned a ton. I was already doing cool UI work with XAML and MVVM but had not worked in a team at such pace, nor seen TeamCity or Continuous Integration with automated tests - over 3,000 - before. Working with a UX team, BAs and exhaustive specs, this was a well run project with 350,000 lines of Silverlight code! It was also fake Agile and burned everyone out.

**July 2008 - January 2011**

## Entrepreneur, Architect & Developer for Myself

*Social Media & Data Analytics Startup*

#### Context and Business Situation

I reworked the "machine learning" algorithm from 2003 (below) and designed and built vuPlan.tv, an app which harvests TV viewing habits data, demographic information, and personalises the customer's programme guide. It offered remote record and synched with Windows Media Center. Sponsored by Microsoft under their BizSpark programme.

#### Engineering Achievements - [See detailed technicals](/me/vuplan-technicals.md)

Designed and coded extreme-performance, sharded in-memory caches, with query self-optimisation and cache routing to speed up mining of millions of rows of data. I query tuned relentlessy and got know the innards of SQL Server. Threat-modelling for DDoS, SQL injection, x-scripting etc. concerns for deployment, reliability and instrumentation on customer computers. Bought and hardened my own public IIS servers and hosted them at a rental facility. The WPF client looked like a phone app, employing MVVM pattern and gestural UI with my own physics math. It dynamically downloaded its ViewModel logic from the cloud at runtime for unobtrusive updates and used full localisation of UI controls and textual elements. I built a top quality Windows Phone 7 application; authored my own controls and made it highly asynchronous and super reliable. During this time I worked directly with the Visual Studio team at Microsoft Redmond to assist with improving the impending VS2010 release.

**On Reflection**

Optimizing a product no one wants for millions of non-existant users is exactly how to fail with a startup.

**August 2001 - July 2008**

## Lead Developer at Société Générale Corporate & Investment Banking

*Banking IT and Large-scale Engineering Ops*

#### Context and Business Situation

Hard to describe. 20 years ago it was unheard of to be a coder who works in a big IT department. Thus, I became an inventor, building tools to automate the industrial scale problems around me and make life easier. I became famous for the stuff I made, it was polished like a boxed app. I've rolled up 7 years of stuff at SGCIB into the following paragraphs.

#### Engineering Achievements - [See detailed technicals](/me/socgen-hector-technicals.md)

Designed a global Wake-on-LAN solution using HTTP APIs as 'proxies' in subnets all around the world, routing wake-requests to machines to control power state. It integrated into other management systems and even Microsoft SMS allowing machines to be woken for critical updates.

Designed a system to plot all the building's physical Cat5 floor-ports on a desk plan, mapped to Cisco switch ports -> ARP -> IP -> hostname -> login -> employee, thus offering location-based services and recommendations for automated software and printer installs.

In 2001 I wrote "FindIP" to amalgamate workstation and HR data providing desktop support with a super user-friendly, drag-and-drop "view of their world". Adoption grew and it became the most-used app in London. I built an Excel add-in offering dozens of cell functions allowing admins to combine live data from the environment with data in disparate systems. With well written F1 help system, it revolutionised the ops team!

Built a system to analyse workstation security group memberships across 14,000 European PCs. Called Logan, backed by a star-schema MSSQL DB and integrated into FindIP and reported via OLAP tools and Excel, it won a global innovation award for its insights in operational risk and impact on Basel II capital adequacy regs. I was appointed innovations "evangelist" for my team.

Won SG Warrants Employee SuperTrader fantasy derivatives trading competition. Won the IT fantasy football league by writing an algorithm.

Lead the design and implementation of a single, standard Windows PC build and automated deployment platform for around 20,000 SGCIB employees, world-wide. With the help of the tools I'd built, my UK team was first to complete with a lead of 17%. This high-profile project encompassed VB and DOS shell scripting, MSI packaging, automation and integration into Windows domain and associated changes to processes.

For a global standardisation project to migrate large Novell Netware user base to NT domains, with hundreds of banking apps to be repackaged, rationalisation of file permissions and network drives, plus BCP provision as a matter of course, I built a collaborative console and auditing tool using VB 6.0 to manage repackaging over 600 applications and analyse the many thousands of file permissions and even automate migration to the new NAS filers. The system proved indispensable and gave IT new insights into managing the desktop estate.

Designed an algorithm that combined PC audit logs with data about the user to disambiguate the software they were using with a confidence score. It learned as you taught it and became completely automated. At the time I didn't realise this was Machine Learning and has a posh term, "Entity Resolution". Years later I used this same algorithm as a learning recommendation engine in my TV viewing habits start-up business.

Writing and supporting VBA, macros and complex Excel workbooks for demanding tradefloor users. 3rd line support for trading apps, common market-data applications, rollouts and connectivity issues.

**On Reflection**

Amazing formative years in a City investment bank but I burned out big time. I was on an uncapped hourly contract and was paid to work as long as I fancied - I almost killed myself. Many of the products I invented and built could be turned into startups. No other bank had this stuff, but we did because we unusually had a developer (me) working and innovating freely on our problems (albeit in secret). I learned that adoption was about making software that solves a real problem and that's a delight to use.

**January 2000 - September 2000**

## Systems Admin at ii.co.uk

*Dotcom Darling*

#### Context and Business Situation

A personal investor and financial website founded by Sherry Coutu CBE.

#### Engineering Achievements

Windows AD and Wintel support. First exposure to real-time price streams and market data, internet backbones and Tier 4 data-centres. Designed and implemented a helpdesk system using Outlook VBA and Exchange forms. Demonstrated a security hole in VBScript which prevented the company from being infected with the ILOVEYOU email worm!!

**On Reflection**

A while ago now but the happiest place I ever worked with the brightest and most creative people and leading edge tech. IPO while I worked there. I learned what a buzzing culture feels like.

**March 1997 - January 2000**

## Product Support Specialist for Microsoft UK

*Microsoft at aged 19*

#### Context and Business Situation

Microsoft's super technical support centre providing assistance for their home and premier corporate customers was, in the UK, run by ICL (and DEC in Ireland). I worked for the former.

#### Engineering Achievements

I held the largest skill-set covering all Windows and DOS versions, all Exchange and Mail products. I was both the youngest NT engineer and the youngest VBA specialist. Learned VB 5.0 and in my own time designed tools for internal MSDN (KB) article authoring and DLL version repository. Won multiple awards for outstanding customer service.

## Credits

I read tons and study technologies and languages just for fun. I listen to many podcasts, I especially like Lenny Ratchitsky's at the moment. Here's just a few of the people's who's ideas I like.

**Childhood**

My childhood was unusual and bears some relevance to today. I've been writing code since I was 8 years old. I had terrible asthma and a computer that didn't run games, but did have a BASIC interpreter built-in, and that meant that a lot of the time I had off school was spent teaching myself to code.

My brother and I were nerdy kids and we used to play with electronics and have competitions on who could design a circuit diagram to solve a particular problem. There were wires all over the house, and even lasers in the back garden where we'd made our own burglar alarm, which actually caught burglars. I once designed a code entry system using relays, buttons and a solenoid. My brother and I ordered the parts and soldered it all together, and it worked! It secured our bedroom cupboard.

Later, I would code a login system for my computer to stop my friends using it, or give them logins. This was using a DOS scripting language called 4DOS. I also made a file encrypter and even figured out a seeded key algorithm to vary the output so repeating characters weren't obvious. This was all before the internet, with no books and very unsupportive parents.

Then I got into computer arts and design. I made nightclub fliers and sold them to promoters. I've even found some of my designs for sale as collectible today! I made them when I was 15. I stopped going to school. Partly because I had become a severe night owl from being up with asthma my whole life, and partly because I'd be coding.

Even though I went to one of the best schools in Britain, it was boring and traditional. I have always been able to teach myself everything, and have always been extremely curious and self motivated - so I found school irrelevant and stopped going. I tried college, but again, I knew everything and I even used to teach some computer arts classes - so that was pointless and I stopped going, too.

When I was hired for Microsoft support, they just wanted smart kids with computing knowledge. Later, I'd have problems getting hired because I was probably in the top 20 most knowledgeable Windows NT engineers in the country, but I was only 21. No recruiter or hiring manager could even begin to ask the right questions - it takes one to know one, and there wasn't anyone.

As an adult, I still struggle. The UK and Europe have stagnated since the birth of the web democratised information and pulled people to Silicon Valley. Here, there are far fewer places to work who understand how to organise highly intelligent, creative and productive people. My CV above is really story of fighting to do great work in old bureaucracies.

**People with ideas**

*   Eric Ries for invalidating assumptions.

*   Kim Scott, Amy Edmondson, Dan North, John Cutler, Jez Humble and Martin Fowler on building together.

*   Ray Dalio, Patty McCord and Max Hastings on meritocracies.

*   John Doerr and Gojko Adžić on impacts and goal alignment.

*   Ohno and Deming for safety and quality.

*   Adam Smith for incentives, inside and out.

*   Kahneman and Tversky for our flawed firmware.

*   Steve Blank and Stephen Wendel on UX.

*   Edward R Tufte and Stephen Few for visually communicating data.

*   Krzysztof Cwalina and Kevlin Henney for "code UX".

*   Randy Shoup for internet scale architectures.

*   Joe Duffy on concurrency and parallelism.
