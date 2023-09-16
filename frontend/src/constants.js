export const jsQuizz = {
    questions: [
      {
        question:
          "Which of the following concepts is not a part of Python?",
        choices: [
          "Loops",
          "Pointers",
          "Dynamic Typing",
          "All of the above",
        ],
        type: "MCQs",
        correctAnswer: "Pointers",
      },
      {
        question: "Which of the following statements are used in Exception Handling in Python?",
        choices: [
          "try",
          "except",
          "finally",
          "All of the above",
        ],
        type: "MCQs",
        correctAnswer: "All of the above",
      },
      {
        question:
          "Which of the following types of loops are not supported in Python?",
        choices: ["for", "while", "do-while", "None of the above"],
        type: "MCQs",
        correctAnswer: "do-while",
      },
      {
        question: "As what datatype are the *args stored, when passed into a function?",
        choices: ["List", "Tuple", "Dictionary", "None of the above"],
        type: "MCQs",
        correctAnswer: "Tuple",
      },
      {
        question: "In which language is Python written?",
        choices: ["C++", "C", "Java", "None of the above"],
        type: "MCQs",
        correctAnswer: "C",
      },
    ],
  };

  export const resultInitialState = {
    score: 0,
    correctAnswers: 0,
    wrongAnswers: 0
  };