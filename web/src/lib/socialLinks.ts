import { GITHUB_URL } from './github.svelte';

// External profile links shared by every surface that shows them
// (Footer, the /signin brand panel) — one place so the URLs can't drift
// the way jobs.company_slug_aliases exists precisely because a fact like
// this once lived in more than one place and disagreed.
export const SOCIAL_LINKS: { provider: string; label: string; href: string }[] = [
  { provider: 'github', label: 'GitHub', href: GITHUB_URL },
];
