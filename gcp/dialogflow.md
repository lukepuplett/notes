# Dialogflow

Lifelike conversational AI with virtual agents.

**Note** - Google also have Contact Center AI, but I'm not sure how exactly it relates yet.

### Key features

- Visual flow builder - as it implies.
- Omnichannel implementation - does _not_ mention Slack or Teams.
- Advanced AI - uses BERT NLU.
- State-based data models - deviate from topic and return.
- End-to-end management - CI/CD, analytics, experiments and bot evaluation.
- For more, see Editions, below.

## Editions

### Dialogflow CX

Advanced agents for large of complex use.

- State-of-the-art BERT NLU model.
- DTMF, live agent handoff, "barge-in", speech timeouts.

### Dialogflow ES (Essentials)

Standard agents for small to medium complexity.

- **No** state-based visualizations.
- It does mention Google Assistant, Slack, Twitter etc.!
- Mentions templates for dining out and hotel booking, navigation, IoT.
- **No** BERT.
- **No** visual builder: form-based builder.
- **No** shared intents, state stuff to switch topics.
- **No** Contact Center AI support.
- Standard simulator.
- **No** test cases for continuous evaluation.
- Basic versions/environments.
- **No** experiments or traffic splitting support.
- **No** user voice ID.

## Notes from video learning materials

### Google's Deconstructing Chatbots tagged YouTube series

https://www.youtube.com/hashtag/deconstructingchatbots

#### Getting Started

The overview schematic shows a few component areas:

- Multi-channel; text chat on web, Facebook, Twitter, Google Assistant on mobile; voice like car, TV and Google Home; phone via gateway.
- Dialogflow Enterprise Edition (old name?)
- Fulfilment; Cloud Function, AppEngine, Compute Engine

Agents are basically the chatbot application; collecting what's being said and mapping it to an intent, taking an action, providing a response. It all starts with a trigger event **utterance**.

"Hey Google, talk to Smart Scheduler."

The phrase "talk to Smart Scheduler" is the _invocation phrase_ and Smart Scheduler is the _invocation name_.

Then the phrase "I want to set up an appointment" contains the intent "set up an appointment".

To control all this, you supply a bunch of example phrases and map them to the supported intents for your bot. Google can then fuzzy match similarly worded phrases to those intents. Then _actions_ and _parameters_ define the variables to collect and store.

"Set an appointment for 5am tomorrow" contains 5am and tomorrow as key information. Those variables are defined as _entities_.

Context is the method for your chatbot to store and access variables to exchange information between intents.

Dialogflow has in-built integration with Cloud Functions and HTTPS endpoints for fulfillment.

#### Appointment Scheduler

https://www.youtube.com/watch?v=oU88sHd6ilE

Create new agent _AppointmentScheduler_ and two default intents are there; welcome and fallback. Create a new intent _Schedule Appointment_ and add some training phrases, "Set an appointment on Wednesday at 2pm" etc.. See that date and time are automatically identified. Then add a response "You're all set for $date at $time, see you then" and then test it in the top right.

Now test it with just "Set an appointment" with no details. To support this, we use _slot filling_; we make the entities as required and it'll make sure to ask for both date and time before responding.

Go to **Actions and parameters** and tick required for the pre-recognised date and time entities. Hit **Define prompts**b and enter some stock ways to ask for that information. Now test it again.

Under **Integrations** on the left, enable the **Web Demo** and use the URL to try out the bot using a basic web page hosted by Google with your users or team.

#### Knowledge Connector Feature

https://www.youtube.com/watch?v=kF33Ime0a2k

Upload an FAQ and have Dialogflow automatically generate questions and answers. Use CSV and HTML files (as of 3 years ago).

Create a Knowledge Connector on the left, name it, choose FAQ, select `text/html` and add the URL of the FAQ on the web. It somehow can pull out the QAs. Click **Add response** to enable the automated responses. Save. Remember to enable it; tick the name of the Knowledge bot and hit **Enable**. Enter some test phrases.

#### Understanding Entities in Dialogflow

https://www.youtube.com/watch?v=UmMRW3kG738

##### System Entities
These are the built-in ones which are sometimes automatically recognised, e.g. @sys.date, time, number, unit-currency, percentage, address, phone-number and email.

##### Developer Entities

You can add your own entities; e.g. add _AppointmentType_ and then add examples of the types with some synonyms like _Drivers License_ as a canonical example and then _license_ and _licence_ and _driving test_. Then _Vehicle Registration_ with like _transfer car_ and _VIN document_ etc. You can have Google automatically append to the list as it figure shit out, and you can upload raw text instead of entering them by hand.

Now when the phrase "Set an appointment for a drivers license at 3pm tomorrow" the term "drivers license" is resolved as the _AppointmentType_ entity. You can now check it as required and it'll be hooked into the slotting filling system.

Remember to add some good prompt examples, and update your response to confirm the type of appointment they made.

##### Session Entities

These are defined at Session ID level, tied to a user and their session with the bot. The presenter oddly glosses over it all and maybe hints that it's done in the API?!

#### Integrating Diialogflow with Actions on Google

https://www.youtube.com/watch?v=z5f52sMgJLQ

One-click integration with Google Assistant, Slack, Facebook Messenger, and Actions on Google which is a platform to extend Google Assistant to smart devices.

Integrations > Explicit invocation (when the "app" is requested by name): Default Welcome Intent.
Integrations > Implicit invocation (when intents trigger "deep link" actions in your app like "OK Google, ask Zipwire to send my timesheet".

**Note** - I'm not sure whether deep linking is to installed phone apps, or is into a bot. Or to a website? 

Manage Assistant App => console.actions.google.com page where you can decide the official public invocation name for your action and how it appears in the Actions Directory, as well as its voice style.

The action/bot needs a description, some sample invocations, branding images (I guess for devices with screens), contact details and privacy and consent.

You can also prepare for deployment in hundres of countries and on phone and speaker "surfaces". There's beta (200 users) and alpha (20 users without Google approval).

#### What is Dialogflow CX?

https://www.youtube.com/watch?v=6_Gilug2QYw

Designed for developers, conversational designers and contact center managers. Visual flow builder for large, complex experiences. Supports multiple flows down authentication, payment processing or orders. Can "call into" different flows. Different organisational departments might control their own flows, handling deviations and transitions. There's environmenting and versioning, too. Many agents under one project.

Native IVR (phone system?) integration. Advanced analytics. Flow design canvas like Visio.

## Notes from Documentation

### Fulfillment

#### Webhook for slot filling

Intents can have required parameters of entities, in which case Dialogflow will keep prompting the user to get information. This is called **slot filling**.

When _webhook for slot filling_ is enabled, a fulfillment webhook request is posted for each conversational turn during slot filling. This lets you further customize the logic during parameter collection. You can customize follow-up questions or set parameter slots from a database.

The webhook payload has `queryResult.parameters` which contains the currently collection ones and those that are unknown are empty strings.

You can set parameters in your response JSON by updating the contexts that are active.

More super useful response JSON examples can be found at the link below, including Slack and Facebook specializations.

https://cloud.google.com/dialogflow/es/docs/fulfillment-webhook#webhook_response
