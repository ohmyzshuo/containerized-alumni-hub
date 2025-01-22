import spacy

def extract_keywords(text, top_n=10):
    """
    Extracts the top N keywords from the given text using spaCy's NLP pipeline.

    Parameters:
    text (str): The input text to extract keywords from.
    top_n (int): The number of top keywords to return.

    Returns:
    list: A list of the top N keywords.
    """
    # Load the English language model
    nlp = spacy.load("en_core_web_sm")

    # Process the text with the NLP pipeline
    doc = nlp(text)

    # Extract the top N ranked phrases
    keywords = [chunk.text.strip() for chunk in doc.noun_chunks]
    keywords = sorted(keywords, key=lambda x: len(x.split()), reverse=True)
    return keywords[:top_n]

# Example usage
sample_text = """
Python is a popular programming language that is widely used for a variety of applications, including web development, data analysis, artificial intelligence, and automation. It is known for its simplicity, readability, and versatility, making it an excellent choice for beginners and experienced developers alike.

One of the key features of Python is its extensive standard library, which provides a wide range of modules and functions for performing common tasks, such as file I/O, network programming, and data manipulation. Additionally, the Python community has developed a vast ecosystem of third-party libraries and frameworks, which extend the functionality of the language and make it suitable for a wide range of use cases.

Python is also known for its strong support for object-oriented programming (OOP), which allows developers to create reusable and maintainable code. It also includes powerful data structures, such as lists, dictionaries, and sets, which make it easy to work with and manipulate data.
"""

keywords = extract_keywords(sample_text, top_n=5)
print(keywords)