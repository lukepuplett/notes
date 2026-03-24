# Luke Puplett - Cloud Product Developer

Hello - I'm a curious, hard-working and collaborative engineer with over two decades experience in very large and very small companies. You'll see a few themes in the stories below; ingenuity, agency and improving how teams build stuff.

I've invented and delivered award-winning tools for managing IT at big banks, built visually rich multi-screen trading apps, tried to build a predictive analytics startup, developed a new app for medics in the US, ran a Python supercomputer and wrote an API auth system that my client patented. I also designed and built my own house, dropped out of one of England's best secondary schools when I was only 13 to write code, and my first job was in Microsoft's tech support team.

Coming from an ops background in vast banking environments where I invented tools for automation and control - SRE before it had a name - I've broad knowledge and interests spanning design and UX to distributed architectures, low-latency trading and lock-free, low-GC code, but also very much company culture and organising people and work.

I work in small batches. I minimise time-to-insight. I automate all the things!! I slow down to speed up. I read voraciously and make time to learn. I teach those around me. I write for others. I focus. I research how best to do something, then I strap in and crack on with it.

This CV is the super long director's cut. If you have a similarly long specification of your role and ideal person, I recommend pasting my CV and your spec into an LLM and asking it if I might be a good or bad fit and why - I think it's the future.

## Experience

**January 2020 - Present Day**

## Founder and Engineer at Zipwire

*Automating Payments, Identity, Blockchain Finance*

#### Context and Business Situation

Zipwire is a platform for recruiters. Zipwire Approve is a worker’s private journal - a WhatsApp bot and a Go CLI for techies - which is used to fill timesheets according to rules so they cannot be wrong. It has workflows for approvals and payouts. Zipwire Collect gathers and scans documents and evidence over the web or WhatsApp, runs ID checks using gov-approved selfie tech. Zipwire Attest can witness these facts on Ethereum.

These apps all feed into the needs of a bigger idea to move invoice factoring onchain and totally automate work, backoffice cashflow and pay, so engineers get paid in minutes from their CLI.

ProofPack is an OSS npm package I created to help folks make selective-disclosure proofs using a verifiable data exchange format. It’s an attested Merkle tree and a JS and .NET SDK to help create, read and verify them. Your bank could use it to let you prove a single line from a statement, without revealing it all.

I started Zipwire during the pandemic and to be at home with my young kids while my wife became a headteacher but I discovered the factoring problem and quickly saw how my work on insurance markets combines with blockchains for a natural solution. There’s a paradox between the truth that most ideas fail because founders give in, and that at some point you actually need to give in.

The timing is off; slow regulation, the bad reputation of crypto plus fragmentation and complexity. My solution hinges on companies broadly issuing employee wallets and using stablecoins to transact.

> "The market may stay irrational longer than you can stay solvent."

#### Engineering Achievements - [See detailed technicals](/me/zipwire-technicals.md)

I learned Google Cloud, React (and SvelteKit), Docker, AI/LLMs, tool/function calling, MCP, even Hardhat and Rust but settled on C# and Golang. It’s a monolithic ASP.NET 6 MPA, Bootstrap, Sass and Knockout.js and some HTMX. I began on Azure using blobs and Cosmos, DevOps/Pipelines before moving it to GCP on Docker in Cloud Run, built on Cloud Build and uses GCS, Firestore and BigQuery.

Zipwire’s fully-featured and way too much to cover here, as it integrates with various LLMs, Google Document AI, OpenAI, Stripe, Twilio, Yoti, FreeAgent, EAS and so on. Everything was built, designed or written by me alone. It uses a modular 'microservices in a monolith' architecture with an in-process event bus.

The aim is to design a DeFi auction house for invoices to solve cashflow problems, so I’ve been studying mechanism design, game theory, Myerson and Milgrom, lots of regs, rewrote half the FCA Handbook. In March 2025 alone I cranked out a 35,000 line Ethereum library, open source, here: [Evoq.Ethereum](https://github.com/luke-puplett/Evoq.Ethereum) with some AI help, but not masses! Proof generation and working with elliptic curve libs.

#### Other

I also hacked together StyleShift using Go, JS and AI models, a Chrome extension which lets people restyle images on property websites - now a built-in feature on RightMove (UK). Some contract work for a small business in my local town, app integration work.

**June 2019 - October 2019**

## Cloud Architect at Incited for Various Insurers

*Big Data Applications and Azure Consultancy*

#### Context and Business Situation

Incited is a startup of experienced Azure devs and data scientists building systems for the UK's big insurers. We worked closely with clients on greenfield cloud-native data solutions around quote stream processing, fraud and realtime analytics.

One client's system ingested up to 50mln long JSON quotes a day from e-commerce websites and looked for patterns in pricing, renewals, churn and deception. We built on Azure PaaS and handed over to the client with training and sometimes help hiring new data talent.

#### Engineering Achievements

I was tasked with redesigning the on-prem, SOAP-based, real-time B2B marketplace that connects many brokers and insurers, Polaris imarket.

I proposed all Azure PaaS, active-active regions, serverless bits and pieces, plus a redesigned stateless JSON API to be entirely hypermedia driven with webhooks. Designed for high-availability and durability of message data in the event of a regional outage, as well as end-to-end security and auth.

I began building components using Azure Event Hubs, Stream Analytics, Service Bus Queues and Functions to determine their suitability and failure modes. You would often find me submitting feedback to the Microsoft docs or engaging with their product teams and even jumping on calls with PMs in Redmond.

**April 2018 - May 2019**

## Senior Systems Analyst and Architect at Equiniti for Canada Life

*Greenfield Pensions and Investments e-Commerce Product Team*

*Link Please read Retirement Advantage below before reading this.*

#### Context and Business Situation

In 2018 the lead on the Retirement Account team invited me to return to the team I had worked with in 2015. The finance team had been sold to Canada Life and Equiniti had bought IT/ops, devs and product, severing product from business - I’d do the opposite.

The 70-strong team had been told they were now working Agile but it was really a waterscrumfall project with a big relaunch months out. I was dropped into the Systems Analysts team to act as the main architect, mentor and provide tech leadership.

The dev Tech Lead and I evangelised a better way using continuous tiny deployments to a production set of services with regular user feedback and instrumentation to guide development, alas, we didn’t win.
The original microservices app was built fast on a "bet the company" basis and its new owner wanted drastic changes. The main Angular quoting app and its backend weren't designed for such change and a previous attempt to add a new product type had been abandoned.

The Tech Lead had run with the things I started in 2016 and had put together a great team of 16 devs. We moved to Azure Service Fabric (like Kubernetes) and integrated the Ops team into Dev. I challenged inertia and advocated to break silos, map impacts and align incentives with the client's goals.

#### Engineering Achievements - [See detailed technicals](/me/equinity-technicals.md)

With a distant launch date and little hope of winning hearts and minds around production pushes, the Tech Lead and I instead solutioned a total replacement of the existing quoting service with a set of new, smaller, event-driven microservices using Rabbit MQ and a new Angular front-end. We built eventually consistent feature stacks - it was a nice design.

I introduced backends-for-frontends to decouple microservices from UX demands, insulating against front-end flux and aiding the transition from service orchestration to event choreography.

I evangelised and mentored 8 analysts and up to 16 developers on designing a real holy-grail, self-documenting HATEOAS application using RESTful APIs the likes of which I have never seen any company produce. I worked on a PoC for a hypermedia client that I've since rewritten and open-sourced as [Surfdude](https://github.com/luke-puplett/Surfdude) - it self-drives hypermedia APIs.

And I designed a SignalR service to push events, including realtime nav to the browser. Introduced Application Insights for observability, though it wasn’t a great product, tbh.

Moved the team to a wiki and went desk to desk teaching non-technical folks Markdown so we could centralise stories, designs and link it all.

Built a PoC for BDD to specify APIs using SpecFlow/Cucumber which drove the app via the hypermedia API via my Surfdude client - seriously cool, pretty sure a world first.

> You've been a massive influence for me professionally and I suspect you're unaware of the number of doors you've opened, and light bulb moments you've provided. This was both in the 2 weeks we initially worked at RA and on a more ongoing basis since you've come back into EQ. I'd like to say thank you and that I'm really, really grateful for all the ways you've helped me.
>
> Technical Lead

**February 2017 - February 2018**

## Technical Lead at Centrica Energy Markets & Trading

*Scaling Python, Energy Options Trading, Valuation & Risk HPC Platform*

#### Context and Business Situation

Within the owner of British Gas sits Centrica EM&T, its prestigious unit that buys, sells and hedges a vast portfolio of energy assets. I ran a small team of devs that worked closely with Quantitative Analytics to value and stress test their nonlinear derivatives portfolios using Python quant models. Our combined systems collated data from LNG cargoes, wind and gas turbines, FX and many other market feeds, then broke the problem into a DAG to run on thousands of compute cores on Azure overnight.

Traders accessed our HPC grid via Excel add-in or an Angular web portal to which they submit and monitor jobs. The on-prem side comprised several Windows services and an Angular SPA backed by an RPC-over-HTTP API. The Azure PaaS side was highly-available, active-active across two environments providing four deployments over two Azure regions so jobs could be moved between them. We were such a large consumer of Azure Batch that I had the personal mobile number of the PM in Redmond since we often discovered regressions in Azure before anyone else.

It could be unreliable and we were often woken at 3am, so I rented a flat near the office. The problems were all downstream of ITIL. A hack culture has built up to get around change control friction from the British Gas business had imposed on what was a quant hedge fund needing to move by the hour! A compounding issue was that the quants whose model would crash on corrupt data were not on pager duty, so they had no incentive to fix their code (and their sleep). If you build it; you run it.


#### Engineering Achievements - [See detailed technicals](/me/centrica-technicals.md)

I oversaw the development and launch of a clever trade version control system that enabled PnL changes to be "explained", like Git Blame for but derivatives trades - not my idea, but it was bloomin’ clever.
Then I began breaking their monoliths and working with a Microsoft MVP on all the CICD automation using VSTS and Azure DevTest Labs, ARM templates, etc.

Introduced Scrum and a huge Kanban board. The company was breaking away from British Gas, moving to Azure, trying to be ‘more Agile’. It took four weeks to get my PC setup with Visual Studio and 9 months to get some whiteboards. Code quality and engineering was a bit too rapid with much hinging around copying files into file shares and cronjobs. The firefighting ate into the time needed to engineer a way out.

I designed much of their MiFID2 reporting solution using Azure Functions, Web Apps, Queues and SQL Data Warehouse with a couple of architects from Microsoft. And for a short while I worked on a greenfield .NET Core 2.0 proper REST API to maintain portfolio stress testing/what-if analysis on Azure App Service and CosmosDB.

Trained at Splunk's London HQ for an observability solution for the distributed HPC environment. The problem here is that Splunk, like any high-perf database, needs custom-spec co-lo servers, but big corp server provisioning tends to be outsourced and a standard VM.

**On Reflection**

The culture was a high-pressure and quite political tradefloor where it was difficult to work deeply on reliable engineering, a common footgun in trading.

**April 2015 - October 2016**

## Senior Developer and Security Expert at Retirement Advantage

*Greenfield Pensions and Investments e-Commerce Product Team*

#### Context and Business Situation

I was initially hired simply to write a SAML 2.0 SSO solution for a flexible new pension product to be used by tens of thousands of public users, but they were in dire need of delivery expertise having never built software like this. They'd been cargo cult software engineering for months, producing Gant charts and piles of Word docs and little code. No process. No source control. No UX. No stories. They had 6 manual testers but no buildable app; they needed help but didn't even know it.

#### Engineering Achievements - [See detailed technicals](/me/mgm-tra-technicals.md)

With the help of a couple of other good contractors we delivered a slick single-page app. Despite having no authority, I steered them towards Git, Visual Studio Online, Slack, a semblance of Scrumban, NuGet, TeamCity, CI and towards the end, PowerShell DSC and even infra-as-code!

The big bang release was a crazy approach in this era so I bought the CEO The Lean Startup. I bought The Phoenix Project for the head of IT and their Head of Solutions Architecture. A Scrum book for the project manager. I tried to convince them of the power of working in small batches and decoupling deploy from release and a continually-deploying set of tested microservices. I tried so hard to explain that the APIs are the app and the BAs should be spec'ing the APIs. I was the only guy that really knew REST, microservices, CICD. It was tough and some developers came in and just left again.

We used Angular on the front-end written mostly by a creative agency, Web API 2.0 on the back with 5 primary microservices. I was officially responsible for security and user-management parts, including encryption, registrations, SSO, AD authentication, IIS and Windows hardening. I built a custom token and API permission and roles system which they patented, and setup a NuGet server. I was eventually allowed to setup a CICD pipeline (!) which was a revelation to all of the product and tech team, though they still insisted on humans manually copying files to production.

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
