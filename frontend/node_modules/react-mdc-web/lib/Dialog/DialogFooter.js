'use strict';

Object.defineProperty(exports, "__esModule", {
  value: true
});

var _propTypes = require('prop-types');

var _propTypes2 = _interopRequireDefault(_propTypes);

var _react = require('react');

var _react2 = _interopRequireDefault(_react);

var _classnames2 = require('classnames');

var _classnames3 = _interopRequireDefault(_classnames2);

var _constants = require('./constants');

function _interopRequireDefault(obj) { return obj && obj.__esModule ? obj : { default: obj }; }

function _defineProperty(obj, key, value) { if (key in obj) { Object.defineProperty(obj, key, { value: value, enumerable: true, configurable: true, writable: true }); } else { obj[key] = value; } return obj; }

var propTypes = {
  className: _propTypes2.default.string,
  children: _propTypes2.default.node
};

var DialogFooter = function DialogFooter(_ref) {
  var className = _ref.className,
      children = _ref.children;

  var last = children.length - 1;
  var actions = _react2.default.Children.map(children, function (action, index) {
    var _classnames;

    var isLastAction = last === index;
    var actionClasses = action.props.className;
    var classes = (0, _classnames3.default)(actionClasses, _constants.FOOTER_BUTTON, (_classnames = {}, _defineProperty(_classnames, _constants.FOOTER_BUTTON_CANCEL, !isLastAction), _defineProperty(_classnames, _constants.FOOTER_BUTTON_ACCEPT, isLastAction), _classnames));
    return _react2.default.cloneElement(action, { key: index, className: classes });
  });

  return _react2.default.createElement(
    'footer',
    {
      className: (0, _classnames3.default)(_constants.FOOTER, className)
    },
    actions
  );
};
DialogFooter.propTypes = propTypes;
exports.default = DialogFooter;