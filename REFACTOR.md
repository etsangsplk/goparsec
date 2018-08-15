



type Parser interface {
    Parse: Scanner -> (Maybe ASTNode, Scanner)
}

Identifier: String

And: Identifier -> Parser... -> Parser
Or: Identifier -> Parser... -> Parser
Kleen: Identifier -> Parser -> Parser
Many: Identifier -> Parser -> Parser
ManyUntil: Identifier -> Parser -> Parser -> Parser
Maybe: Identifier -> Parser -> Parser

Shortcircuit: Parser -> (ASTNode -> Scanner -> Maybe ASTNode) -> Parser // enables lookahead and more

Char: Identifier -> Parser
Float: Identifier -> Parser
Hex: Identifier -> Parser
Int: Identifier -> Parser
Oct: Identifier -> Parser
String: Identifier -> Parser
Ident: Identifier -> Parser
 
Atom: Identifier -> String -> Parser
AtomExact: Identifier -> String -> Parser
 
Regex: String

Token: Identifier -> Regex -> Parser
TokenExact: Identifier -> Regex -> Parser 
OrToken: Identifier -> Regex... -> Parser

End: -> Parser
NoEnd: -> Parser

Source {
    File: String
    Line: Int
    Index: Int
}

ASTNode {
    Name: Identifier
    GetChildren: -> []ASTNode
    GetSegment: -> String // substring responsible for this node
    GetSource -> Source
    IsTerminal: -> Bool
}

NodeTransformer: ASTNode -> Maybe ASTNode // because transforming into nothing is a valid thing

ASTNodePredicate: ASTNode -> Bool

TransformAll: ASTNode -> NodeTransformer -> Maybe ASTNode // Note this is is a bottom up recursive transform
TransformMatch: ASTNode -> ASTNodePredicate -> NodeTransformer -> Maybe ASTNode
TransformMatches: ASTNode -> [](ASTNodePredicate -> NodeTransformer) -> Maybe ASTNode

Out: {}interface

NodeGenerator: ASTNode -> [](Out) -> Out // second arguments is generated children. Terminals have this empty.
GenerateAll: ASTNode -> NodeGenerator -> Out
FoldUp: ASTNode -> [](ASTNodePredicate, Generator) -> Out


IsA: Identifier -> ASTNodePredicate
//       IsA("FORLOOP")(node) -> T/F

Map: []ASTNode -> (ASTNode -> Out) -> []Out

QueryFilter: ASTNode -> ASTNodePredicate -> []ASTNode
Terminals: ASTNode -> []ASTNode
QueryMany: ASTNode -> QueryString -> []ASTNode
QueryOne: ASTNode -> QueryString -> Maybe ASTNode

