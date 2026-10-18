# Publishing to the Chrome Web Store

Everything the dashboard asks for, kept next to the manifest it has to match.
Reviewers cross-check this copy against `extension/wxt.config.ts` and the code —
when one changes, change the other.

## Release checklist

1. **Bump the version** in `extension/package.json` (`version`). The store
   rejects a re-upload of a version it already has.
2. **Get the package** — either `cd extension && npm run zip`
   (`.output/hireme-extension-<version>-chrome.zip`) or download the
   `chrome-mv3` artifact from the `build` workflow.
3. **Upload** to the [Developer Dashboard](https://chrome.google.com/webstore/devconsole)
   (one-time $5 registration + a verified contact email).
4. Fill the listing, privacy practices, and permission justifications below.
5. Submit. Expect an in-depth manual review — `<all_urls>` guarantees it.
6. **After the item exists**, take its extension ID from the dashboard and add
   `https://<id>.chromiumapp.org/` to the redirect-URI allowlist of
   `hiremetoday.ca/api/v1/auth/extension/connect`. Until that lands, sign-in fails
   for everyone who installs from the store — the published build has no
   `key`, so its ID differs from the pinned dev one.

Still manual, not in the repo: the privacy-policy page on `hiremetoday.ca` and
1–5 screenshots at 1280×800 (side panel over a real job posting).

## Listing

**Name:** hireme

**Short description** (≤132 chars):

> A job-application agent in your side panel: it reads the posting you are on,
> matches it to your CV, and fills the form.

**Single purpose** (the store requires exactly one):

> Assist the user with job applications on the page they are currently viewing.

**Detailed description:**

> hireme puts a job-application agent in Chrome's side panel, next to whatever
> page you are on.
>
> - Ask it about the posting you are reading. It reads the page itself when your
>   question needs it, so you can just say "is this a fit?" and get an answer
>   about the tab in front of you.
> - See how the posting matches your hireme profile — a deterministic
>   skill-coverage card, not a guess.
> - Let it fill the application form. It reads the form's fields and writes the
>   values from your hireme profile; you review and submit yourself.
>
> You need a hireme account. Sign in once from the panel; the extension does
> nothing until you do.

**Category:** Productivity. **Language:** English.

## Permission justifications

Paste one per field in the dashboard.

| Permission | Justification |
| --- | --- |
| `storage` | Stores the hireme session token issued by sign-in, plus nothing else. Local to the browser profile. |
| `tabs` | The panel needs the active tab's id and URL to know which page the user is asking about and to route the agent's read/fill calls to it. |
| `sidePanel` | The entire UI is a side panel opened from the toolbar icon. |
| `scripting` | Injects the reader into the active tab's frames on demand, so the agent can read the application form and fill the fields the user asked it to. |
| `activeTab` | Scopes page access to the tab the user explicitly acted on from the panel. |
| `identity` | `launchWebAuthFlow` runs "Sign in with hireme" and returns the session token to the extension. Not used for Google account data. |
| `host_permissions: <all_urls>` | Applications live on arbitrary hosts — Greenhouse, Lever, Workday, Ashby, and thousands of company career pages. The user decides which page the agent works on, so the host set cannot be enumerated in advance. Page content is read only while the panel is open and only in service of what the user asked for — the question they typed, or the Autofill they pressed. Nothing is read in the background: the channel the reads travel over exists only for as long as the panel is. |

## Privacy practices

Data collected — declare all three:

- **Personally identifiable information** — name, email, phone and CV fields from
  the user's hireme profile, fetched to fill application forms.
- **Authentication information** — the hireme session JWT, kept in
  `chrome.storage.local`.
- **Website content** — the URL, title and visible text (capped at 5000
  characters, see `extension/lib/scraper.ts`) of the page open while the user is
  asking about it, plus the form fields Autofill reads. Sent to `hiremetoday.ca`,
  which is the only host this extension contacts.

  Two things worth stating plainly, because they changed: the agent now decides
  when to read the page — during a turn, when the user's question is about what is
  in front of them — rather than the user pressing a button to hand it over. And a
  page read this way is kept in that conversation's transcript, which the user can
  read and delete from their hireme account.

  Reads are bounded and visible. Only ordinary web pages are read: a browser
  settings page, another extension's page, or a local file is refused before it is
  touched. And every read is named in the conversation itself — the panel shows
  which page it read, without the query string, so a read the user did not expect
  is one they can see.

Not collected: health, financial, location, personal communications, browsing
history, keystroke logging.

Certifications to tick: data is not sold to third parties, not used or
transferred for purposes unrelated to the extension's single purpose, and not
used for creditworthiness or lending.

The privacy policy at the URL you enter must say the same thing and state
compliance with the Chrome Web Store limited-use requirements. It has to be
reachable over HTTPS without a login — the review bot crawls it.

## Automated publishing (after the first manual release)

The Chrome Web Store API only updates an **existing** item — it cannot create the
first listing, so everything above stays a one-time manual pass. Once that first
release is live, `.github/workflows/extension.yml` uploads and submits every
subsequent version bump automatically (gated on the GitHub Release step actually
having created a new tag, so an unchanged version is never re-uploaded — the
store rejects that anyway).

It's a no-op until three secrets exist on the repo. One-time setup:

1. [Google Cloud Console](https://console.cloud.google.com/) → new (or existing)
   project → **APIs & Services → Library** → enable the **Chrome Web Store API**.
2. **APIs & Services → Credentials** → **Create credentials → OAuth client ID**
   → Application type **Desktop app**. Note the client id and client secret.
3. Get a refresh token once, from a machine with a browser (this is the only
   interactive step): run
   ```
   npx --yes chrome-webstore-upload-cli@3 refresh-token \
     --client-id <CLIENT_ID> --client-secret <CLIENT_SECRET>
   ```
   sign in with the **same Google account that owns the Developer Dashboard
   item**, and copy the refresh token it prints.
4. Add three repo secrets (Settings → Secrets and variables → Actions):
   `CHROME_WEBSTORE_CLIENT_ID`, `CHROME_WEBSTORE_CLIENT_SECRET`,
   `CHROME_WEBSTORE_REFRESH_TOKEN`.

`CHROME_WEBSTORE_EXTENSION_ID` is not a secret — it's hardcoded in the workflow
as the published item's id (`ijfaechijopdlikalojadpojmpilplnj`, from the
dashboard URL), which is **not** the same id `wxt.config.ts`'s pinned dev key
produces locally.

Submitting a build this way still enters Google's normal manual review queue —
`<all_urls>` guarantees that regardless of who uploads it. Automation removes
the manual dashboard upload step, not the wait.
