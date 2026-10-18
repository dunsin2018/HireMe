import { defineMessages } from '$lib/i18n/t';

// "Gmail" and "HireMe" are brand names and stay as-is in every locale.
export const messages = defineMessages(
  {
    connected: 'Connected',
    gmailHint: 'Pull replies from your own Gmail (needs Google sign-in).',
    notAvailable: 'Not available yet.',
    manageIntegrations: 'Manage in Integrations',
    connectIntegrations: 'Connect in Integrations',
    hiremeMailboxLabel: 'hireme mailbox',
    copyAddressTitle: 'Copy address',
    mailboxHint: 'Use this address when you apply — replies land here.',
    release: 'Release',
    creating: 'Creating…',
    getMailbox: 'Get a hireme mailbox',
    getMailboxHint: 'Get an address on our domain — no Google needed.',
    releaseDialogTitle: 'Release your hireme mailbox?',
    releaseDialogDescription: 'Its received mail is deleted.',
    claimFailed: 'Failed to create a mailbox.',
    releaseFailed: 'Failed to release the mailbox.',
  },
  {
    ru: {
      connected: 'Подключено',
      gmailHint: 'Получайте ответы из своего Gmail (нужен вход через Google).',
      notAvailable: 'Пока не доступно.',
      manageIntegrations: 'Управлять в Интеграциях',
      connectIntegrations: 'Подключить в Интеграциях',
      hiremeMailboxLabel: 'Почтовый ящик hireme',
      copyAddressTitle: 'Скопировать адрес',
      mailboxHint: 'Используйте этот адрес при отклике — ответы будут приходить сюда.',
      release: 'Отказаться',
      creating: 'Создаём…',
      getMailbox: 'Получить почтовый ящик hireme',
      getMailboxHint: 'Получите адрес на нашем домене — без Google.',
      releaseDialogTitle: 'Отказаться от почтового ящика hireme?',
      releaseDialogDescription: 'Полученные письма будут удалены.',
      claimFailed: 'Не удалось создать почтовый ящик.',
      releaseFailed: 'Не удалось отказаться от почтового ящика.',
    },
  },
);
