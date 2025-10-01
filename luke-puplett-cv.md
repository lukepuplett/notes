# Luke Puplett - Cloud Product Developer

I'm a curious, hard-working and collaborative engineer with over two decades experience in very large and very small companies. You'll see a few themes in the stories below; inventiveness, ingenuity and pushing for a better way.

I've invented and built award-winning tools for managing IT at big banks, built visually rich multi-screen trading apps, failed as an entrepreneur in predictive analytics, way back in 2003 I made up a machine learning algo and solved a gnarly problem, developed a new app for medics in the US, ran a Python supercomputer, and wrote an API auth system that my client patented. I also designed and built my own house. I dropped out of one of England's best secondary schools when I was only 13 to write code. My first job was in Microsoft's tech support team.

I've an unusually broad knowledge and interests. I'm from an ops background in vast banking environments where I built tools for automating everything, however I could have gone down a UX path. I can design visually—I even oil paint—and I can also design internet-scale architectures. On a micro scale, I have written lock-free, low-contention, low-GC code for trading.

I work in small batches. I minimise time-to-insight. I automate all the things. I slow down to speed up. I read voraciously and make time to learn. I teach those around me. I write. I focus. I research how best to do something, then I strap in and crack on with it.

This version of my CV is very long. Please skim or read enough to form a decision. Else, paste this into an LLM and ask me questions (that's probably the future of hiring, anyway).

## Experience

**January 2020 - Present Day**

## Founder and Engineer at Zipwire

*Timesheets and Payments Side Hustle*

#### Context and Business Situation

Zipwire is a platform for recruiters. 'Zipire Approve' is for time journalling; it combines WhatsApp and AI so you can keep a diary using natural speech, it has an approval workflow and a backoffice and payments workflow. 'Zipwire Collect', pairs WhatsApp and AI to collect documents from people and run ID checks using a government approved selfie AI model.

I started Zipwire to get through the pandemic, to be at home with a young family while my wife became a headteacher. I didn't expect it to consume my life. I learned Google Cloud, Docker, AI/LLMs, even Hardhat and Rust but settled on C# and Golang. It's a monolithic ASP.NET 6 MPA, Bootstrap, Sass and Knockout.js. I began on Azure using blobs and Cosmos, DevOps/Pipelines before I moved to GCP, containerised in Cloud Run, built on Cloud Build and uses blobs, Firestore, and Big Query.

#### Engineering Achievements

[See detailed technicals](/me/zipwire-technicals.md)

Zipwire's fully-featured and too much to cover here, as it integrates with various LLMs, Google Document AI, OpenAI TTS, Stripe, Twilio and Yoti. Everything was built, designed or written by me alone. It uses a modular 'microservices in a monolith' architecture which I discuss here.

I'll spare you the details but the aim is to integrate with the Ethereum blockchain. In March alone I cranked out a 35,000 line Ethereum library, open source, here: [Evoq.Ethereum](https://github.com/luke-puplett/Evoq.Ethereum) with the help of Cursor.ai.

**June 2019 - October 2019**

## Cloud Architect at Incited for Various Insurers

*Big Data Applications and Azure Consultancy*

#### Context and Business Situation

Incited is a startup of experienced Azure devs and data scientists building systems for the UK's big insurers. We worked closely with clients on greenfield cloud-native data solutions around quote stream processing, fraud and realtime analytics.

One client's system ingested up to 50mln long JSON quotes a day from e-commerce websites and looked for patterns in pricing, renewals, churn and deception. We built on Azure PaaS and handed over to the client with training and sometimes help hiring new data talent.

#### Engineering Achievements

I was tasked with redesigning the on-prem, SOAP-based marketplace that sits beneath much of the UK insurance industry, Polaris imarket. Think how Go Compare gets quotes from all the insurers.

I proposed all Azure PaaS, active-active regions, serverless bits and pieces, plus a redesigned stateless JSON API to be entirely hypermedia driven with webhooks. Designed for high-availability and durability of message data in the event of a regional outage, as well as end-to-end security and auth.

I began building components using Azure Event Hubs, Stream Analytics, Service Bus Queues and Functions to determine their suitability and failure modes. You can often find me submitting feedback to the Microsoft docs or engaging with their product teams and even jumping on calls with PMs in Redmond.

**April 2018 - May 2019**

## Senior Systems Analyst and Architect at Equiniti for Canada Life

*Greenfield Pensions and Investments e-Commerce Product Team*

*Link Please read Retirement Advantage below before reading this.*

#### Context and Business Situation

In 2018 the Tech Lead on the Retirement Account team asked me to return to the team I had worked with in 2015. The business had been sold to Canada Life and sadly the former CEO had severed the product people from the business, selling the product team to Equiniti who are charged back to Canada Life.

The 70-strong team were "told" they were working Agile but it was waterscrumfall with a delivery milestone months out. I was dropped into the Systems Analysts team to act as the main architect, mentor and provide thought leadership.

The Tech Lead and I strenuously evangelised a better way using continuous deployments to production with regular feedback and instrumentation to guide development. As in 2015, changing their ways was super challenging.

The original microservices app was built fast, on a "bet the company" basis. Now Canada Life wanted drastic changes. The main Angular quoting app and its backend were'nt designed for such change and a previous attempt to add a new product type had been abandoned.

The Tech Lead had ran with the things I started in 2016 and had put together a great development team of 16. We moved to Azure and incorporated Ops. I challenged inertia and pushed hard to break silos, map impacts and align incentives with the client's goals, which was a struggle since the business had been outsourced/disintegrated.

#### Engineering Achievements

[See detailed technicals](/me/equinity-technicals.md)

With a far out deployment date and small hope of winning hearts and minds around weekly production pushes, the Tech Lead and I instead solutioned a complete replacement of the existing quoting service with a set of new, smaller, event-driven microservices using Rabbit MQ and a new Angular front-end. The idea was to build loosely-coupled service stacks behind the product's main feature areas.

I introduced backends-for-frontends to decouple microservices from UX demands helping insulate against front-end change and also aid the transition from service orchestration to event choreography.

I evangelised and mentored 8 analysts and up to 16 developers on designing a real holy-grail, self-documenting HATEOAS application using RESTful APIs the likes of which I have never seen any company produce. I worked on a PoC for a hypermedia client that I've since rewritten as an open-source library called [Surfdude](https://github.com/luke-puplett/Surfdude).

It was eventually consistent so I designed a SignalR service to push updates and even nav to the browser. Pitched and won use of Azure Service Fabric. Introduced Application Insights for observability, though I regret it.

Moved the product team to a wiki and went desk to desk teaching non-technical folks Markdown. They used Office docs in a SharePoint accessed via RDP, ugh. Information used to be scattered across thousands of files, Jiras and emails.

Built a PoC for BDD to specify APIs using SpecFlow which drove the app via the hypermedia API, via my Surfdude client. It was very cool.

> You've been a massive influence for me professionally and I suspect you're unaware of the number of doors you've opened, and light bulb moments you've provided. This was both in the 2 weeks we initially worked at RA and on a more ongoing basis since you've come back into EQ. I'd like to say thank you and that I'm really, really grateful for all the ways you've helped me.
>
> Technical Lead

**February 2017 - February 2018**

## Technical Lead at Centrica Energy Markets & Trading

*Scaling Python, Energy Options Trading, Valuation & Risk HPC Platform*

#### Context and Business Situation

Within Centrica—the owner of British Gas—is EM&T, a prestigious unit that buys, sells and hedges options in energy markets, like a commodities hedge fund. I had a small team of banking developers that worked closely with Quantitative Analytics to value Centrica's nonlinear derivatives portfolios using a Python quant model, scaled to collate data, break the problem up, and run on thousands of compute cores on Azure overnight, with monitoring tools.

Traders access our HPC grid via Excel add-in or an Angular web portal to which they submit and monitor jobs. The on-prem side comprises several Windows Services and an Angular SPA backed by an RPC-over-HTTP API. The Azure PaaS side is highly-available, active-active with 2 environments providing 4 deployments over two Azure regions and jobs can be moved between them. We were such a large consumer of Azure Batch that I had the personal mobile number of the PM in Redmond since we often discovered regressions in Azure before anyone else.

There were many problems, most of which could be traced back to working around the friction imposed by ITIL. It could have been an excellent system and a joy to work on, but it had been hacked together in response to rapidly evolving needs.

#### Engineering Achievements

[See detailed technicals](/me/centrica-technicals.md)

I oversaw the development and launch of a clever trade version control system that enabled PnL changes to be "explained", like Git Blame for but derivatives trades.

Then I began breaking their monoliths and working with a Microsoft MVP on all the CICD automation using VSTS and Azure DevTest Labs, ARM templates, etc.

Introduced Scrum and a huge Kanban board. The company were in two worlds, Azure coming in and talk of Agile and DevOps, while trying to break free of old British Gas. It took four weeks to get my PC setup with Visual Studio and 9 months to get some whiteboards! Code quality and engineering was rapid MVP quality with much designed around dropping files into file shares and scheduled tasks. They had severe reliability issues and little time to focus on proper engineering between fighting fires.

I designed much of their MiFID2 reporting solution using Azure Functions, Web Apps, Queues and SQL Data Warehouse with a couple of architects from Microsoft. And for a short while I worked on a greenfield .NET Core 2.0 proper REST API to maintain portfolio stress testing/what-if analysis on Azure App Service and CosmosDB.

**On Reflection**

The culture was difficult; a stereotypical, high-pressure tradefloor where tactical fixes and continual fire-fighting had become the MO. It was difficult to slow down or even work deeply on engineering reliable solutions.

**April 2015 - October 2016**

## Senior Developer and Security Expert at Retirement Advantage

*Greenfield Pensions and Investments e-Commerce Product Team*

#### Context and Business Situation

I was initially hired simply to write a SAML 2.0 SSO solution for a flexible new pension product to be used by tens of thousands of public users, but they were in dire need of delivery expertise having never built software like this. They'd been cargo cult software engineering for months, producing Gant charts and piles of Word docs and little code. No process. No source control. No UX. No stories. They had 6 manual testers but no buildable app; they needed help but didn't even know it.

#### Engineering Achievements

[See detailed technicals](/me/mgm-tra-technicals.md)

With the help of a couple of other good contractors we managed to deliver a pretty slick single-page app. Despite having zero authority and in a low-trust setting, I steered them towards Git, Visual Studio Online, Slack, a semblance of Scrumban, NuGet, TeamCity, CI and towards the end, PowerShell DSC and even infra-as-code.

The big bang release was a crazy approach in this era so I bought the CEO The Lean Startup. I bought The Phoenix Project for the head of IT and their Head of Solutions Architecture. A Scrum book for the project manager. Without power, it was like having the Curse of Cassandra as I fought to convince them to work in small batches and decouple deploy from release. They would slip back to old ways, struggling to comprehend a continually-deploying set of tested microservices. I tried so hard to explain that the APIs are the app and the BAs should be spec'ing the APIs. I was the only guy that really knew how to do REST, microservices, CICD. It was tough and some developers came in and just left again.

We used Angular on the front-end written mostly by a creative agency, Web API 2.0 on the back with 5 primary microservices. I was mostly responsible for the security and user-management parts, including encryption, registrations, SSO, AD authentication, IIS and Windows hardening. I built a custom token and API permission and roles system which they patented, and setup NuGet for dependency/library management. I was eventually allowed to setup a CICD pipeline which was, no surprise, a revelation to all of the product and tech team, though they insisted on using humans to manually push to production!

MSSQL with EF, MongoDB, and the app interfaced with an off-shelf pensions BO product via SOAP. The most interesting technical take-away from this contract was the importance of a very considered API design and even only using GET and POST and less strict web APIs for the UI to minimise the client-side code and, frankly, appease the JavaScript developers, a pattern we now call Backends for Frontends.

I was kept on for an extra month when their Solvency II team failed a regulatory deadline, so I automated and CI'ed complex SSIS/MSSQL using PowerShell+TeamCity.

**On Reflection**

As of December 2023 the app has been rebuilt three time but the original auth service I wrote is still going strong. It's a cool thing and could be a business on its own, like Okta. My proudest achievements were the patent but also that two pen-testing companies found nothing, and by the end, the SAs and BAs knew Postman inside out :). In honesty, I was quite disruptive, I had to be, but I changed them for the better.

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

#### Engineering Achievements

[See detailed technicals](/me/vuplan-technicals.md)

Designed and coded extreme-performance, sharded in-memory caches, with query self-optimisation and cache routing to speed up mining of millions of rows of data. I query tuned relentlessy and got know the innards of SQL Server. Threat-modelling for DDoS, SQL injection, x-scripting etc. concerns for deployment, reliability and instrumentation on customer computers. Bought and hardened my own public IIS servers and hosted them at a rental facility. The WPF client looked like a phone app, employing MVVM pattern and gestural UI with my own physics math. It dynamically downloaded its ViewModel logic from the cloud at runtime for unobtrusive updates and used full localisation of UI controls and textual elements. I built a top quality Windows Phone 7 application; authored my own controls and made it highly asynchronous and super reliable. During this time I worked directly with the Visual Studio team at Microsoft Redmond to assist with improving the impending VS2010 release.

**On Reflection**

Optimizing a product no one wants for millions of non-existant users is exactly how to fail with a startup.

**August 2001 - July 2008**

## Lead Developer at Société Générale Corporate & Investment Banking

*Banking IT and Large-scale Engineering Ops*

#### Context and Business Situation

Hard to describe. 20 years ago it was unheard of to be a coder who works in a big IT department. Thus, I became an inventor, building tools to automate the industrial scale problems around me and make life easier. I became famous for the stuff I made, it was polished like a boxed app. I've rolled up 7 years of stuff at SGCIB into the following paragraphs.

#### Engineering Achievements

[See detailed technicals](/me/socgen-hector-technicals.md)

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

Amazing formative years in a City investment bank, but I burned out. Big time. Much of the stuff I built could be productized and sold. No other bank had this stuff. We had it because we unusually had a developer (me) working and innovating freely in the team. I learned that adoption was about making software that's a delight to use.

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
