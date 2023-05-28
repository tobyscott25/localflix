export function formatDate(dateString: string, format: 'short' | 'long') {
	const date = new Date(dateString)

	const shortFormatOptions: Intl.DateTimeFormatOptions = {
		day: 'numeric',
		month: 'numeric',
		year: 'numeric',
		hour: 'numeric',
		minute: 'numeric',
	}

	const longFormatOptions: Intl.DateTimeFormatOptions = {
		weekday: 'long',
		day: 'numeric',
		month: 'long',
		year: 'numeric',
		hour: 'numeric',
		minute: 'numeric',
	}

	const formattedDate = date.toLocaleString(
		'en-AU',
		format === 'short' ? shortFormatOptions : longFormatOptions
	)

	return formattedDate
}
