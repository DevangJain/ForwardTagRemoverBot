/*  Copyright (C) 2020 by Anandpskerala@Github, < https://github.com/Anandpskerala >.
 *
 * This file is part of < https://github.com/Anandpskerala/ForwardTagRemoverBot > project,
 * and is released under the "GNU v3.0 License Agreement".
 * Please see < https://github.com/Anandpskerala/ForwardTagRemoverBot/blob/master/LICENSE >
 *
 * All rights reserved.
 */

package commands

import (
	"github.com/PaulSonOfLars/gotgbot"
	"github.com/PaulSonOfLars/gotgbot/ext"
	"github.com/PaulSonOfLars/gotgbot/parsemode"
	"go.uber.org/zap"
)

func Help(b ext.Bot, u *gotgbot.Update) error {

	helpButton := [][]ext.InlineKeyboardButton{make([]ext.InlineKeyboardButton, 2), make([]ext.InlineKeyboardButton, 1)}

	helpButton[0][0] = ext.InlineKeyboardButton{
		Text: "Support Channel 🌐",
		Url:  "https://t.me/Devajebots",
	}

	helpButton[0][1] = ext.InlineKeyboardButton{
		Text: "My Creater 🔧",
		Url:  "https://t.me/Devaje",
	}

	helpButton[1][0] = ext.InlineKeyboardButton{
		Text: "Report a Bug 🔧",
		Url:  "https://t.me/Devajegroup",
	}

	markup := ext.InlineKeyboardMarkup{InlineKeyboard: &helpButton}

	msg := b.NewSendableMessage(u.EffectiveChat.Id, "Forward me any File,Video,Audio,Photo or Anything And \nI will Send You the File Back\n\n`How to Set Caption?`\nReply your Caption to a File,Photo,Audio,Media\nIf You want any Support Join Our Support Group 😊")
	msg.ReplyToMessageId = u.EffectiveMessage.MessageId
	msg.ReplyMarkup = &markup
	msg.ParseMode = parsemode.Markdown
	_, err := msg.Send()
	if err != nil {
		b.Logger.Warnw("Error in sending", zap.Error(err))
	}
	return err
}
