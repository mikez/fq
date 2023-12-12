# -*- coding: utf-8 -*-
# Make mail data: text/plain and text/html.
# Usage: python3 make_mail.py

from email.charset import Charset, add_charset
import email.encoders
from email.mime.image import MIMEImage
from email.mime.multipart import MIMEMultipart
from email.mime.text import MIMEText
import io

from PIL import Image, ImageDraw, ImageFont  # pip install pillow


CHARSETS = ["utf-8", "iso-8859-1", "iso-8859-15", "us-ascii", "windows-1252"]
CONTENT_TRANSFER_ENCODINGS = ["quoted-printable", "base64", "none"]


def main():
    languages = make_languages()

    # single part
    for language in languages:
        for charset in language["charsets"]:
            for cte in language["cte"]:  # Content-Transfer-Encoding
                message = make_singlepart_message(language, charset, cte)
                save_bytes(
                    bytes(message), f"email_single_{language['id']}_{charset}_{cte}.eml"
                )

    # multipart alternative with plaintext and HTML
    language, charset, cte = languages[0], "utf-8", "quoted-printable"
    message = make_multipart_alternative_message(language, charset, cte)
    save_bytes(bytes(message), f"email_multi_{language['id']}_{charset}_{cte}.eml")

    # multipart with plaintext and images interleaved
    language, charset, cte = languages[0], "us-ascii", "quoted-printable"
    message = make_image_message(language, charset, cte)
    save_bytes(bytes(message), f"email_image_{language['id']}_{charset}_{cte}.eml")


def make_languages():
    english = {
        "id": "en",
        "subject": "Test Email",
        "text": (
            "Greetings, fellow traveler! "
            "This is a test email with {mediatype} wrapped using the MIME format "
            "specified in RFC 2045 of 1996."
        ),
        "plain": "plain text content",
        "html": "HTML content",
        "charsets": ["us-ascii"],
        "cte": ["quoted-printable", "base64", "none"],
    }

    swedish = {
        "id": "sv",
        "subject": "Test mejl",
        "text": "Hej! Det här är ett mejl med {mediatype}.",
        "plain": "textinnehåll",
        "html": "HTML-innehåll",
        "charsets": ["utf-8", "iso-8859-1"],
        "cte": ["quoted-printable", "none"],
    }

    german = {
        "id": "de",
        "subject": "Test E-Mail",
        "text": "Hallo! Dies ist eine E-Mail mit {mediatype} und dem Eurozeichen: €.",
        "plain": "Textinhalt",
        "html": "HTML-Inhalt",
        "charsets": ["utf-8", "iso-8859-15", "windows-1252"],
        "cte": ["quoted-printable"],
    }

    emoji = {
        "id": "emoji",
        "subject": "Test 📧",
        "text": "👋\n\nThis is a 📧 with {mediatype} and emojis 🙃.",
        "plain": english["plain"],
        "html": english["html"],
        "charsets": ["utf-8"],
        "cte": ["quoted-printable"],
    }

    return [english, swedish, german, emoji]


def make_singlepart_message(language_data, charset, cte):
    set_email_library_charset_default(charset, cte)

    text = language_data["text"].format(mediatype=language_data["plain"])
    result = MIMEText(text, "plain", charset)

    # Add header
    result["From"] = "sender@example.com"
    result["To"] = "recipient@example.com"
    result["Subject"] = language_data["subject"]

    return result


def make_multipart_alternative_message(language_data, charset, cte):
    result = MIMEMultipart("alternative")
    set_email_library_charset_default(charset, cte)

    # Add header
    result["From"] = "sender@example.com"
    result["To"] = "recipient@example.com"
    result["Subject"] = language_data["subject"]

    # Add plain text part
    text = language_data["text"].format(mediatype=language_data["plain"])
    plain_part = MIMEText(text, "plain", charset)
    result.attach(plain_part)

    # Add HTML part
    text = language_data["text"].format(mediatype=language_data["html"])
    html_text = f"<html><body><p>{text}</p></body></html>"
    html_part = MIMEText(html_text, "html", charset)
    result.attach(html_part)

    return result


def make_image_message(language_data, charset, cte):
    result = MIMEMultipart()
    set_email_library_charset_default(charset, cte)

    # Add header
    result["From"] = "sender@example.com"
    result["To"] = "recipient@example.com"
    result["Subject"] = language_data["subject"]

    # Add plain text part
    text = language_data["text"].format(mediatype=language_data["plain"])
    plain_part = MIMEText(text + "\n\n", "plain", charset)
    result.attach(plain_part)

    # Add Image part
    image_bytes = create_image("fq", 200, 100)
    image_part = MIMEImage(image_bytes, "image/png", name="image.png")
    image_part["Content-Disposition"] = "inline"
    result.attach(image_part)

    # Add another plain text part
    plain_part = MIMEText("\n\nGoodbye!", "plain", charset)
    result.attach(plain_part)

    return result


# Helpers


def create_image(text, width, height, fontsize=40):
    image = Image.new("RGB", (width, height), "#001f3f")
    draw = ImageDraw.Draw(image)
    font = ImageFont.load_default(size=fontsize)
    bbox = draw.textbbox((0, 0), text, font=font)
    text_position = ((width - bbox[2]) // 2, (height - bbox[3]) // 2)
    draw.text(text_position, text, font=font, fill="#ffdc00")
    result = io.BytesIO()
    image.save(result, format="png")
    return result.getvalue()


def exclude(iterable, *target_items):
    return [item for item in iterable if item not in target_items]


def set_email_library_charset_default(charset, cte):
    """
    Change `email.charset.CHARSETS`.

    In the `email` library, charsets have hard-coded defaults set for
    the Content-Transfer-Encodings header for a given charset;
    change the default option here.

    See https://github.com/python/cpython/blob/main/Lib/email/charset.py
    """
    header_enc, body_enc, output_conv = email.charset.CHARSETS[charset]
    body_enc = {
        "quoted-printable": email.charset.QP,
        "base64": email.charset.BASE64,
        "none": None,
    }[cte]
    if not output_conv:
        output_conv = charset
    add_charset(charset, header_enc, body_enc, output_conv)


def save_bytes(bytestring, filepath):
    with open(filepath, "wb") as file:
        file.write(bytestring)


if __name__ == "__main__":
    main()
